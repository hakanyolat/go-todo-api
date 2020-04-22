package model

import (
	"github.com/jinzhu/gorm"
	"regexp"
)

// Migrate table(s)
func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Task{})
	return db
}

// For migration tests
func GetMigrationQueries() map[interface{}][]string {
	return map[interface{}][]string{
		Task{}: []string{
			regexp.QuoteMeta(`CREATE TABLE "tasks" ("id" INTEGER AUTO_INCREMENT,"created_at" TIMESTAMP,"updated_at" TIMESTAMP,"deleted_at" TIMESTAMP,"title" VARCHAR(255),"completed" BOOLEAN , PRIMARY KEY ("id"))`),
			regexp.QuoteMeta(`CREATE INDEX idx_tasks_deleted_at ON "tasks"(deleted_at) `),
		},
	}
}