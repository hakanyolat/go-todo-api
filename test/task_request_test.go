package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hakanyolat/go-todo-api/model"
	"github.com/hakanyolat/go-todo-api/model/request"
	"net/http"
	"reflect"
	"testing"
)

func TestTask_CreateGetRequest(t *testing.T) {
	want := &request.GetTaskRequest{
		Id: uint64(1),
	}
	wantBytes, _ := json.Marshal(want)

	hr, _ := http.NewRequest(http.MethodGet, testConfig.TCP.Host, nil)
	hr = mux.SetURLVars(hr, map[string]string{
		"id": fmt.Sprint(want.Id),
	})

	got, _ := request.NewGetTaskRequest(hr)
	gotBytes, _ := json.Marshal(got)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %s but got %s", wantBytes, gotBytes)
	}
}

func TestTask_CreateTaskRequest(t *testing.T) {
	want := &request.CreateTaskRequest{
		Title:     "Test",
		Completed: true,
	}

	jsonBytes, _ := json.Marshal(want)

	hr, _ := http.NewRequest(http.MethodPost, testConfig.TCP.Host, bytes.NewReader(jsonBytes))
	got, _ := request.NewCreateTaskRequest(hr)
	gotBytes, _ := json.Marshal(got)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %s but got %s", jsonBytes, gotBytes)
	}
}

func TestTask_UpdateTaskRequest(t *testing.T) {
	want := &request.UpdateTaskRequest{
		Id:        uint64(1),
		Completed: true,
	}

	jsonBytes, _ := json.Marshal(want)

	wantBody := &request.UpdateTaskRequest{
		Completed: want.Completed,
	}

	wantBodyBytes, _ := json.Marshal(wantBody)

	hr, _ := http.NewRequest(http.MethodPut, testConfig.TCP.Host, bytes.NewReader(wantBodyBytes))
	hr = mux.SetURLVars(hr, map[string]string{
		"id": fmt.Sprint(want.Id),
	})

	got, _ := request.NewUpdateTaskRequest(&model.Task{}, hr)
	gotBytes, _ := json.Marshal(got)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %s but got %s", jsonBytes, gotBytes)
	}
}