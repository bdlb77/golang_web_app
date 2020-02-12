package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"golang_web_app/src/github.com/bdlb77/webapp/model"
	"golang_web_app/src/github.com/bdlb77/webapp/viewmodel"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
}

func (s shop) RegisterRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	http.HandleFunc("/shop/", s.handleShop)
	// /shop/ -> Go will recognize this as a param route
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		categoryId, _ := strconv.Atoi(matches[1])
		s.handleCategory(w, r, categoryId)
	} else {
		categories := model.GetCategories()
		vm := viewmodel.NewShop(categories)
		s.shopTemplate.Execute(w, vm)
	}
}
func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryId int) {
	products := model.GetProductsForCategory(categoryId)
	vm := viewmodel.NewShopDetail(products)
	fmt.Printf("my vm: %v", vm.Products)
	s.categoryTemplate.Execute(w, vm)

}
