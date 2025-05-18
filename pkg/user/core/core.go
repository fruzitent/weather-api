package core

import (
	_ "git.fruzit.pp.ua/weather/api/pkg/user/core/command"
	_ "git.fruzit.pp.ua/weather/api/pkg/user/core/query"
)

type App struct {
	Command Command
	Query   Query
}

type Command struct{}

type Query struct{}
