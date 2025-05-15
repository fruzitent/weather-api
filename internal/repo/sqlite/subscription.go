package sqlite

import "database/sql"

type subscription struct {
	db *sql.DB
}

func NewSubscriptionRepo(db *sql.DB) *subscription {
	return &subscription{db}
}
