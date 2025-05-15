package sqlite

import (
	"database/sql"

	"git.fruzit.pp.ua/weather/api/internal/repo"
)

type subscription struct {
	db *sql.DB
}

func NewSubscriptionRepo(db *sql.DB) repo.ISubscription {
	return &subscription{db}
}
