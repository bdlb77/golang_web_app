package viewmodel

import (
	"golang_web_app/src/github.com/bdlb77/webapp/model"
)

type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

func NewShopDetail(products []model.Product) ShopDetail {
	result := ShopDetail{
		Title:    "Lemonade Stand Supply",
		Active:   "shop",
		Products: []Product{},
	}
	for _, p := range products {
		result.Products = append(result.Products, ProductToVM(&p))
	}
	return result
}
