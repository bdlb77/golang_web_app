package controller

import (
	"html/template"
	"net/http"

	"golang_web_app/src/github.com/bdlb77/webapp/viewmodel"
)

type home struct {
	homeTemplate *template.Template
}

// startup func creates route handlers then registers the routes
func (h home) RegisterRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
}

// handlers handleFunc interface.
func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}
