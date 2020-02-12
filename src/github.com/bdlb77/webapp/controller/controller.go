package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	shopController shop
)

func StartUp(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	shopController.shopTemplate = templates["shop.html"]
	homeController.RegisterRoutes()
	shopController.RegisterRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
