package submission

import (
	"context"
	"fmt"

	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	mainDB         *pgxpool.Pool
	submissionPool map[int]*models.Submission
}

func NewRepository(mainDB *pgxpool.Pool) *repository {
	return &repository{
		mainDB:         mainDB,
		submissionPool: make(map[int]*models.Submission),
	}
}

func (r *repository) CreateSubmission(
	owner string,
	problemId string,
	content string,
	languageId string,
) (submission models.Submission, err error) {
	ctx := context.Background()

	verdict := "Pending"
	row := r.mainDB.QueryRow(
		ctx,
		`INSERT INTO Submission
		(owner, content, language_id, verdict, problem_id)
		VALUES
		($1, $2, $3, $4, $5)
		RETURNING id, submitted_at;`,
		owner,
		content,
		languageId,
		verdict,
		problemId,
	)
	if err = row.Scan(&submission.Id, &submission.SubmittedAt); err != nil {
		err = fmt.Errorf("unable to scan row: %s", err.Error())
		return
	}
	submission.Owner = owner
	submission.Content = content
	submission.LanguageId = languageId
	submission.ProblemId = problemId

	r.submissionPool[submission.Id] = &submission
	return
}

func (r *repository) UpdateSubmissionVerdictInPool(id int, verdict string) (submission *models.Submission, found bool) {
	submission, found = r.submissionPool[id]
	if !found {
		return
	}
	submission.Verdict = verdict
	return
}

func (r *repository) GetSubmissionInPool(id int) (submission *models.Submission, found bool) {
	submission, found = r.submissionPool[id]
	return
}
