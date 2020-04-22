package service

import (
	"github.com/hakanyolat/go-todo-api/app"
	"github.com/hakanyolat/go-todo-api/model/response"
	"net/http"
)

type WelcomeService struct {
	app.Service
}

func (s *WelcomeService) Provide() {
	s.Router.Get("/", s.GetHome)
}

func (s *WelcomeService) GetHome(w http.ResponseWriter, r *http.Request) {
	s.SendResponse(w, response.NewMessageResponse("Welcome"), http.StatusOK)
}
