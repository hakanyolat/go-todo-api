package service

import (
	"github.com/hakanyolat/go-todo-api/app"
	"github.com/hakanyolat/go-todo-api/model"
	"github.com/hakanyolat/go-todo-api/model/request"
	"github.com/hakanyolat/go-todo-api/model/response"
	"github.com/hakanyolat/go-todo-api/repository"
	"net/http"
)

type TaskService struct {
	app.Service
}

func (s *TaskService) Provide() {
	s.Router.Get("/tasks", s.GetTasks)
	s.Router.Post("/tasks", s.CreateTask)
	s.Router.Get("/tasks/{id:[0-9]+}", s.GetTask)
	s.Router.Put("/tasks/{id:[0-9]+}", s.UpdateTask)
	s.Router.Delete("/tasks/{id:[0-9]+}", s.DeleteTask)
}

func (s *TaskService) GetTasks(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewTaskRepository(s.DB)

	if tasks, err := repo.FindAll(); err != nil {
		s.SendErrorResponse(w, &app.HttpError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	} else {
		s.SendResponse(w, response.NewTasksResponse(tasks), http.StatusOK)
	}
}

func (s *TaskService) GetTask(w http.ResponseWriter, r *http.Request) {
	var err error
	var req *request.GetTaskRequest
	var httpErr *app.HttpError

	if req, httpErr = request.NewGetTaskRequest(r); httpErr != nil {
		s.SendErrorResponse(w, httpErr)
		return
	}

	repo := repository.NewTaskRepository(s.DB)
	var task *model.Task

	if task, err = repo.Find(req.Id); err != nil {
		s.SendErrorResponse(w, &app.HttpError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	} else {
		s.SendResponse(w, response.NewTaskResponse(*task), http.StatusOK)
	}
}

func (s *TaskService) CreateTask(w http.ResponseWriter, r *http.Request) {
	var err error
	var req *request.CreateTaskRequest
	var httpErr *app.HttpError

	if req, httpErr = request.NewCreateTaskRequest(r); httpErr != nil {
		s.SendErrorResponse(w, httpErr)
		return
	}

	repo := repository.NewTaskRepository(s.DB)
	var task *model.Task

	if task, err = repo.Create(req); err != nil {
		s.SendErrorResponse(w, &app.HttpError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	s.SendResponse(w, response.NewTaskResponse(*task), http.StatusOK)
}

func (s *TaskService) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var err error
	var req *request.UpdateTaskRequest
	var httpErr *app.HttpError
	var task *model.Task

	repo := repository.NewTaskRepository(s.DB)

	var getReq *request.GetTaskRequest
	if getReq, httpErr = request.NewGetTaskRequest(r); httpErr != nil{
		s.SendErrorResponse(w, httpErr)
		return
	}

	if task, err = repo.Find(getReq.Id); err != nil{
		s.SendErrorResponse(w, &app.HttpError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	if req, httpErr = request.NewUpdateTaskRequest(task, r); httpErr != nil {
		s.SendErrorResponse(w, httpErr)
		return
	}

	if task, err = repo.Update(req); err != nil {
		s.SendErrorResponse(w, &app.HttpError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	s.SendResponse(w, response.NewTaskResponse(*task), http.StatusOK)
}

func (s *TaskService) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var err error
	var httpErr *app.HttpError
	var task *model.Task
	var req *request.GetTaskRequest

	repo := repository.NewTaskRepository(s.DB)
	result := false

	if req, httpErr = request.NewGetTaskRequest(r); httpErr != nil{
		s.SendErrorResponse(w, httpErr)
		return
	}

	if task, err = repo.Find(req.Id); err != nil{
		s.SendErrorResponse(w, &app.HttpError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	if result, err = repo.Delete(task); err != nil {
		s.SendErrorResponse(w, &app.HttpError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	s.SendResponse(w, result, http.StatusNoContent)
}
