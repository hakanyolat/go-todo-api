package model

import (
	"github.com/hakanyolat/go-todo-api/app"
)

var Registry = []app.ModelInterface{
	&Task{},
}
