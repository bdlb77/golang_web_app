package controller

import (
	"fmt"
	"html/template"
	"log"
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
	// otherwise b/c of GZIP.. if content type is not set to text
	// it will download as a GZIP
	w.Header().Add("Content-Type", "text/html")
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error Logging:  %v \n", err))
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		if email == "hi@gmail.com" && password == "password" {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		} else {
			vm.Email = email
			vm.Password = password
		}
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}
