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

var testConfig config.Config
var testApp *app.App
var Mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	testConfig = config.Get(string(app.Env.Mock))

	_, Mock, _ = sqlmock.NewWithDSN(testConfig.DB.GetDSN())

	testApp = app.NewApp().Mock()
	testApp.Configure(testConfig)
	testApp.Init(model.Registry, service.Registry)
	os.Exit(m.Run())
}
