package auth

import (
	"context"
	"fmt"

	"github.com/FreeJ1nG/cpduel-backend/app/models"
	"github.com/georgysavva/scany/v2/pgxscan"
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

func (r *repository) CreateUser(username string, fullName string, passwordHash string) (user models.User, err error) {
	ctx := context.Background()

	_, err = r.mainDB.Exec(
		ctx,
		`INSERT INTO Account (username, full_name, password_hash) VALUES ($1, $2, $3);`,
		username,
		fullName,
		passwordHash,
	)

	if err != nil {
		return
	}

	user = models.User{
		Username:     username,
		FullName:     fullName,
		PasswordHash: passwordHash,
	}
	return
}

func (r *repository) GetUserByUsername(username string) (user models.User, err error) {
	ctx := context.Background()

	err = pgxscan.Get(
		ctx,
		r.mainDB,
		&user,
		`SELECT * FROM Account WHERE username = $1;`,
		username,
	)

	if err != nil {
		err = fmt.Errorf("unable to get user by username: %s", err.Error())
		return
	}
	return
}
