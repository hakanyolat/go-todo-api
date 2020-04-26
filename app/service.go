package app

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
)

type ServiceInterface interface {
	Init(r *Router, db *gorm.DB)
	Provide()
}

type Service struct {
	Router *Router
	DB     *gorm.DB
}

func (s *Service) Init(r *Router, db *gorm.DB) {
	s.Router = r
	s.DB = db
}

func (s *Service) SendResponse(w http.ResponseWriter, v interface{}, status int) {
	response, err := json.Marshal(v)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

func (s *Service) SendErrorResponse(w http.ResponseWriter, e *HttpError) {
	s.SendResponse(w, map[string]interface{}{"message": e.Message}, e.Code)
}
