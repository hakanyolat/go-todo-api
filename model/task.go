package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Task model
type Task struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}