package controller

import (
	"html/template"
	"net/http"

	"golang_web_app/src/github.com/bdlb77/webapp/viewmodel"
)

type home struct {
	homeTemplate  *template.Template
	loginTemplate *template.Template
}

// startup func creates route handlers then registers the routes
func (h home) RegisterRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

// handlers handleFunc interface.
func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	h.loginTemplate.Execute(w, vm)
}
