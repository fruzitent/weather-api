package sqlite

import "database/sql"

type weather struct {
	db *sql.DB
}

func NewWeatherRepo(db *sql.DB) *weather {
	return &weather{db}
}
