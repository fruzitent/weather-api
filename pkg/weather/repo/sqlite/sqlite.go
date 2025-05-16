package sqlite

import (
	"database/sql"

	"git.fruzit.pp.ua/weather/api/pkg/weather/repo"
)

type weather struct {
	db *sql.DB
}

func New(db *sql.DB) repo.IWeather {
	return &weather{db}
}
