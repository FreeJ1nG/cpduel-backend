package submission

import "github.com/jackc/pgx/v5/pgxpool"

type repository struct {
	mainDB *pgxpool.Pool
}

type Repository interface{}
