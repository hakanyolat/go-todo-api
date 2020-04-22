package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	app *App
	mux *mux.Router
}

func (r *Router) Get(route string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc(route, fn).Methods("GET")
}

func (r *Router) Post(route string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc(route, fn).Methods("POST")
}

func (r *Router) Head(route string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc(route, fn).Methods("HEAD")
}

func (r *Router) Put(route string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc(route, fn).Methods("PUT")
}

func (r *Router) Patch(route string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc(route, fn).Methods("PATCH")
}

func (r *Router) Options(route string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc(route, fn).Methods("OPTIONS")
}

func (r *Router) Delete(route string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc(route, fn).Methods("DELETE")
}

func (r *Router) SetNotFoundHandler(fn func(w http.ResponseWriter, r *http.Request)) {
	r.mux.NotFoundHandler = http.HandlerFunc(fn)
}

func NewRouter(app *App) *Router {
	return &Router{
		app: app,
		mux: mux.NewRouter(),
	}
}
