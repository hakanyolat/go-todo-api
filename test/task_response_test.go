package test

import (
	"encoding/json"
	"github.com/hakanyolat/go-todo-api/model"
	"github.com/hakanyolat/go-todo-api/model/response"
	"github.com/jinzhu/gorm"
	"reflect"
	"testing"
)

func TestTask_TaskResponse(t *testing.T) {
	want := response.TaskResponse{
		ID:        1,
		Title:     "Test",
		Completed: false,
	}

	task := model.Task{
		Model: gorm.Model{
			ID: 1,
		},
		Title:     "Test",
		Completed: false,
	}

	got := response.NewTaskResponse(task)

	if !reflect.DeepEqual(want, got) {
		wantBytes, _ := json.Marshal(want)
		gotBytes, _ := json.Marshal(got)
		t.Errorf("Expected %s but got %s", wantBytes, gotBytes)
	}
}

func TestTask_TasksResponse(t *testing.T) {
	want := []response.TaskResponse{
		{
			ID:        1,
			Title:     "Test",
			Completed: false,
		},
		{
			ID:        2,
			Title:     "Test 2",
			Completed: true,
		},
	}

	tasks := []model.Task{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Title:     "Test",
			Completed: false,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Title:     "Test 2",
			Completed: true,
		},
	}

	got := response.NewTasksResponse(tasks)

	if !reflect.DeepEqual(want, got) {
		wantBytes, _ := json.Marshal(want)
		gotBytes, _ := json.Marshal(got)
		t.Errorf("Expected %s but got %s", wantBytes, gotBytes)
	}
}
