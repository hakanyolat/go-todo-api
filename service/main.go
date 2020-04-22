package service

import "github.com/hakanyolat/go-todo-api/app"

var Registry = []app.ServiceInterface{
	&DefaultService{},
	&WelcomeService{},
	&TaskService{},
}
