package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang_web_app/src/github.com/bdlb77/webapp/model"
	"golang_web_app/src/github.com/bdlb77/webapp/viewmodel"
)

type home struct {
	homeTemplate   *template.Template
	loginTemplate  *template.Template
	signUpTemplate *template.Template
}

// startup func creates route handlers then registers the routes
func (h home) RegisterRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
	http.HandleFunc("/sign_up", h.handleSignUp)

}

// handlers handleFunc interface.
func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	// otherwise b/c of GZIP.. if content type is not set to text
	// it will download as a GZIP
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

		// write DB Query and Update for Time Now
		if user, err := model.Login(email, password); err == nil {
			log.Printf("user has Logged in: %v \n", user.Email)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		} else {
			log.Printf("failed to log in user: %v. Error: %v", user.Email, err)
			vm.Email = email
			vm.Password = password
		}
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}

func (h home) handleSignUp(w http.ResponseWriter, r *http.Request) {
	// receive vm for SignUp
	vm := viewmodel.NewSignUp()
	// get form data
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error for Sign Up:  %v \n", err))
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		FirstName := r.Form.Get("first_name")
		LastName := r.Form.Get("last_name")

		// send to model to handle user create
		err = model.SignUp(email, password, FirstName, LastName)

		// log in user.
		// user, err := model.Login(email, password)
		if err == nil {
			// if not err, redirect to home
			log.Printf("%v Has signed up! \n", email)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		} else {
			// else print issue signing up and set fields back to what they typed.
			log.Printf("failed to log in user: %v. Error: %v", email, err)
			vm.Email = email
			vm.FirstName = FirstName
			vm.LastName = LastName
			vm.Password = password
		}
	}
	// Execute signup template. with response and vm
	h.signUpTemplate.Execute(w, vm)
}
