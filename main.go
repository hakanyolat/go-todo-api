package main

import (
	"github.com/hakanyolat/go-todo-api/app"
	"github.com/hakanyolat/go-todo-api/config"
	"github.com/hakanyolat/go-todo-api/model"
	"github.com/hakanyolat/go-todo-api/service"
)


func main() {
	env := app.Env.Default
	conf := config.Get(env)

	api := app.NewApp()
	api.Configure(conf)
	api.Init(model.Registry, service.Registry)
	api.Run()
}
