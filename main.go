package main

import (
	"github.com/hakanyolat/go-todo-api/app"
	"github.com/hakanyolat/go-todo-api/config"
	"github.com/hakanyolat/go-todo-api/service"
)


func main() {
	configuration := config.GetConfig()
	application := app.NewApp()

	application.Init(configuration)
	application.Register(service.Registry...)
	application.Run()
}
