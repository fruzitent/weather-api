package sqlite

import (
	"database/sql"

	"git.fruzit.pp.ua/weather/api/internal/repo"
)

type weather struct {
	db *sql.DB
}

func NewWeatherRepo(db *sql.DB) repo.IWeather {
	return &weather{db}
}
