package core

import (
	_ "git.fruzit.pp.ua/weather/api/pkg/weather/core/command"
	"git.fruzit.pp.ua/weather/api/pkg/weather/core/query"
)

type App struct {
	Command Command
	Query   Query
}

type Command struct{}

type Query struct {
	Current query.CurrentHandler
}
