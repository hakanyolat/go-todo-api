package main

import (
	"flag"
	"fmt"
	"github.com/hakanyolat/go-todo-api/app"
	"github.com/hakanyolat/go-todo-api/config"
	"github.com/hakanyolat/go-todo-api/model"
	"github.com/hakanyolat/go-todo-api/service"
	"strings"
)

var Env string
var Conf config.Config

func init() {
	flag.StringVar(&Env, "env", string(app.Env.Default), fmt.Sprintf("The environment option for application. Possible values: %s.", strings.Join(app.Env.GetPossibleValues(), ", ")))
	flag.Parse()
	Conf = config.Get(Env)
}

func main() {
	api := app.NewApp()
	api.Configure(Conf)
	api.Init(model.Registry, service.Registry)
	api.Run()
}
