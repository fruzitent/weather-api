package sqlite

import (
	"database/sql"

	"git.fruzit.pp.ua/weather/api/pkg/weather/repo"
)

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) repo.IRepo {
	return &Repo{db}
}
