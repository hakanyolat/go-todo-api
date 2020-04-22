package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hakanyolat/go-todo-api/model"
	"github.com/hakanyolat/go-todo-api/model/request"
	"github.com/hakanyolat/go-todo-api/repository"
	"regexp"
	"testing"
)

func TestTaskRepository_Find(t *testing.T) {
	repo := repository.NewTaskRepository(testApp.GetDB())

	id := uint64(1)
	cols := []string{"id"}
	rows := sqlmock.NewRows(cols)
	query := regexp.QuoteMeta(`
		SELECT * FROM "tasks"  
			WHERE "tasks"."deleted_at" IS NULL AND (("tasks"."id" = 1)) 
			ORDER BY "tasks"."id" ASC 
		LIMIT 1`)
	Mock.ExpectQuery(query).WillReturnRows(rows)

	_, _ = repo.Find(id)

	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestTaskRepository_FindAll(t *testing.T) {
	repo := repository.NewTaskRepository(testApp.GetDB())

	cols := []string{"id"}
	rows := sqlmock.NewRows(cols)
	query := regexp.QuoteMeta(`SELECT * FROM "tasks"  WHERE "tasks"."deleted_at" IS NULL`)
	Mock.ExpectQuery(query).WillReturnRows(rows)

	_, _ = repo.FindAll()

	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestTaskRepository_Create(t *testing.T) {
	repo := repository.NewTaskRepository(testApp.GetDB())

	Mock.ExpectBegin()
	Mock.ExpectExec("INSERT INTO \"tasks\".*").WillReturnResult(sqlmock.NewResult(0, 0))
	Mock.ExpectCommit()

	_, _ = repo.Create(&request.CreateTaskRequest{
		Title:     "Test",
		Completed: false,
	})

	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestTaskRepository_Update(t *testing.T) {
	repo := repository.NewTaskRepository(testApp.GetDB())

	Mock.ExpectBegin()
	Mock.ExpectExec("UPDATE \"tasks\".*").WillReturnResult(sqlmock.NewResult(0, 0))
	Mock.ExpectCommit()

	_, _ = repo.Update(&request.UpdateTaskRequest{
		Id:        5,
		Completed: true,
		Task: &model.Task{
			Title:     "Test",
			Completed: true,
		},
	})

	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestTaskRepository_Delete(t *testing.T) {
	repo := repository.NewTaskRepository(testApp.GetDB())

	// Soft delete...
	Mock.ExpectBegin()
	Mock.ExpectExec("UPDATE \"tasks\".*").WillReturnResult(sqlmock.NewResult(0, 0))
	Mock.ExpectCommit()

	_, _ = repo.Delete(&model.Task{
		Title:     "Test",
		Completed: true,
	})

	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
