package service

import (
	"github.com/hakanyolat/go-todo-api/app"
	"github.com/hakanyolat/go-todo-api/model/response"
	"net/http"
)

type DefaultService struct {
	app.Service
}

func (s *DefaultService) Provide() {
	s.Router.SetNotFoundHandler(s.NotFoundHandler)
}

func (s *DefaultService) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	s.SendResponse(w, response.NewMessageResponse("Endpoint was not found"), http.StatusNotFound)
}
