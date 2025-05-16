package sqlite

import (
	"database/sql"

	"git.fruzit.pp.ua/weather/api/pkg/subscription/repo"
)

type subscription struct {
	db *sql.DB
}

func New(db *sql.DB) repo.ISubscription {
	return &subscription{db}
}
