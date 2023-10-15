package problem

import (
	"context"
	"fmt"

	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	mainDB *pgxpool.Pool
}

func NewRepository(mainDB *pgxpool.Pool) *repository {
	return &repository{
		mainDB: mainDB,
	}
}

func (r *repository) CreateProblem(
	problemId string,
	title string,
	timeLimit string,
	memoryLimit string,
	inputType string,
	outputType string,
	difficulty int,
	body string,
) (problem models.Problem, err error) {
	ctx := context.Background()

	_, err = r.mainDB.Exec(
		ctx,
		`INSERT INTO Problem 
		(id, title, time_limit, memory_limit, input_type, output_type, difficulty, body)
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id)
		DO NOTHING;`,
		problemId,
		title,
		timeLimit,
		memoryLimit,
		inputType,
		outputType,
		difficulty,
		body,
	)

	if err != nil {
		err = fmt.Errorf("failed to create problem: %s", err.Error())
		return
	}

	problem = models.Problem{
		Id:          problemId,
		Title:       title,
		TimeLimit:   timeLimit,
		MemoryLimit: memoryLimit,
		InputType:   inputType,
		OutputType:  outputType,
		Difficulty:  difficulty,
		Body:        body,
	}
	return
}

func (r *repository) CreateLanguages(languages []models.Language) (err error) {
	ctx := context.Background()
	rows := make([][]interface{}, 0)
	for _, language := range languages {
		langInterface := []interface{}{language.Id, language.Value}
		rows = append(rows, langInterface)
	}
	_, err = r.mainDB.CopyFrom(
		ctx,
		pgx.Identifier{"language"},
		[]string{"id", "value"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return
	}
	return
}

func (r *repository) CreateProblemLanguages(languages []models.Language, problemId string) (err error) {
	ctx := context.Background()
	rows := make([][]interface{}, 0)
	for _, language := range languages {
		problemLanguageInterface := []interface{}{problemId, language.Id}
		rows = append(rows, problemLanguageInterface)
	}
	_, err = r.mainDB.CopyFrom(
		ctx,
		pgx.Identifier{"languageofproblem"},
		[]string{"problem_id", "language_id"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return
	}
	return
}

func (r *repository) GetProblemById(problemId string) (problem models.Problem, err error) {
	ctx := context.Background()
	err = pgxscan.Get(
		ctx,
		r.mainDB,
		&problem,
		`SELECT * FROM Problem WHERE id = $1;`,
		problemId,
	)
	if err != nil {
		err = fmt.Errorf("failed to fetch problem from the database: %s", err.Error())
		return
	}
	return
}

func (r *repository) GetLanguageWithIds(languageIds []string) (languages []models.Language, foundIds map[string]bool, err error) {
	ctx := context.Background()

	rows, err := r.mainDB.Query(
		ctx,
		`SELECT * FROM Language WHERE id = ANY($1::VARCHAR(40)[]);`,
		languageIds,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	foundIds = make(map[string]bool)

	for rows.Next() {
		var language models.Language
		err = rows.Scan(&language.Id, &language.Value)
		if err != nil {
			return
		}
		languages = append(languages, language)
		foundIds[language.Id] = true
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

func (r *repository) GetLanguagesOfProblemById(problemId string) (languages []models.Language, err error) {
	ctx := context.Background()
	rows, err := r.mainDB.Query(
		ctx,
		`SELECT L.* FROM LANGUAGE L
		INNER JOIN LanguageOfProblem LP ON L.id = LP.language_id
		WHERE LP.problem_id = $1`,
		problemId,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	languages = []models.Language{}

	for rows.Next() {
		var language models.Language
		err = rows.Scan(&language.Id, &language.Value)
		if err != nil {
			return
		}
		languages = append(languages, language)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

func (r *repository) DeleteProblemById(problemId string) (err error) {
	ctx := context.Background()
	_, err = r.mainDB.Exec(
		ctx,
		`DELETE FROM Problem WHERE id = $1;`,
		problemId,
	)
	if err != nil {
		return
	}
	return
}
