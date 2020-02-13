package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController         home
	shopController         shop
	standLocatorController standLocator
)

func StartUp(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	homeController.loginTemplate = templates["login.html"]
	standLocatorController.standLocatorTemplate = templates["stand_locator.html"]
	shopController.categoryTemplate = templates["shop_details.html"]
	shopController.shopTemplate = templates["shop.html"]

	homeController.RegisterRoutes()
	standLocatorController.RegisterRoutes()
	shopController.RegisterRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
