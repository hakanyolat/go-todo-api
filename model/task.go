package model

import (
	"github.com/hakanyolat/go-todo-api/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Task struct {
	app.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t *Task) Migrate(db *gorm.DB) {
	db.AutoMigrate(&t)
}

func(t *Task) WillSeed() []app.ModelInterface {
	return []app.ModelInterface{
		&Task{
			Title:     "test 1",
			Completed: true,
		},
		&Task{
			Title:     "test 2",
			Completed: false,
		},
		&Task{
			Title:     "test 3",
			Completed: true,
		},
	}
}