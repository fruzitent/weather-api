package core

import (
	"git.fruzit.pp.ua/weather/api/pkg/user/core/command"
	_ "git.fruzit.pp.ua/weather/api/pkg/user/core/query"
)

type App struct {
	Command Command
	Query   Query
}

type Command struct {
	Subscribe command.SubscribeHandler
}

type Query struct{}
