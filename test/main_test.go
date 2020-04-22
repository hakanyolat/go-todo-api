package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hakanyolat/go-todo-api/app"
	"github.com/hakanyolat/go-todo-api/config"
	"github.com/hakanyolat/go-todo-api/model"
	"github.com/hakanyolat/go-todo-api/service"
	"os"
	"testing"
)

var testConfig *config.Config
var testApp *app.App
var Mock sqlmock.Sqlmock

func ExpectMigrations() {
	for _, v := range model.GetMigrationQueries() {
		for _, q := range v{
			Mock.ExpectExec(q).WillReturnResult(sqlmock.NewResult(0,0))
		}
	}
}

func TestMain(m *testing.M) {
	testConfig = &config.Config{
		TCP: &config.TCPConfig{
			Host: "127.0.0.1",
			Port: "8080",
		},
		DB: &config.DBConfig{
			Dialect:  "sqlmock",
			Host:     "sqlmock",
			Port:     "3306",
			Username: "mock",
			Password: "mock",
			Name:     "mock",
			Charset:  "utf8",
		},
	}

	_, Mock, _ = sqlmock.NewWithDSN(testConfig.DB.GetDSN())

	ExpectMigrations()

	testApp = app.NewApp()
	testApp.Init(testConfig)
	testApp.Register(service.Registry...)

	os.Exit(m.Run())
}
