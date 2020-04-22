package request

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hakanyolat/go-todo-api/app"
	"github.com/hakanyolat/go-todo-api/model"
	"net/http"
	"strconv"
)

type GetTaskRequest struct {
	Id uint64
}

type CreateTaskRequest struct {
	Title     string
	Completed bool
}

type UpdateTaskRequest struct {
	Id        uint64
	Completed bool
	Task      *model.Task
}

func NewGetTaskRequest(r *http.Request) (*GetTaskRequest, *app.HttpError) {
	params := mux.Vars(r)
	var id uint64
	var err error

	if _, ok := params["id"]; ok {
		id, err = strconv.ParseUint(params["id"], 10, 64)

		if err != nil {
			return nil, &app.HttpError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		if id == 0 {
			return nil, &app.HttpError{
				Code:    http.StatusBadRequest,
				Message: "\"id\" must be greater than zero",
			}
		}
	} else {
		return nil, &app.HttpError{
			Code:    http.StatusBadRequest,
			Message: "\"id\" is required",
		}
	}

	return &GetTaskRequest{Id: id}, nil
}

func NewCreateTaskRequest(r *http.Request) (*CreateTaskRequest, *app.HttpError) {
	req := CreateTaskRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, &app.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	defer r.Body.Close()

	if req.Title == "" {
		return nil, &app.HttpError{
			Code:    http.StatusBadRequest,
			Message: "\"title\" is required",
		}
	}
	return &req, nil
}

func NewUpdateTaskRequest(task *model.Task, r *http.Request) (*UpdateTaskRequest, *app.HttpError) {
	req := UpdateTaskRequest{
		Task: task,
	}

	params := mux.Vars(r)
	var id uint64

	if _, ok := params["id"]; ok {
		var err error
		id, err = strconv.ParseUint(params["id"], 10, 64)

		if err != nil {
			return nil, &app.HttpError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		if id == 0 {
			return nil, &app.HttpError{
				Code:    http.StatusBadRequest,
				Message: "\"id\" must be greater than zero",
			}
		}
	} else {
		return nil, &app.HttpError{
			Code:    http.StatusBadRequest,
			Message: "\"id\" is required",
		}
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, &app.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	defer r.Body.Close()

	req.Id = id
	return &req, nil
}