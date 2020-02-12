package model

import "fmt"

type Product struct {
	Name             string
	DescriptionShort string
	DescriptionLong  string
	PricePerLiter    float32
	PricePer10Liter  float32
	Origin           string
	IsOrganic        bool
	ImageURL         string
	Id               int
	CategoryId       int
}

func GetProductsForCategory(categoryId int) []Product {
	result := []Product{}
	for _, prod := range products {
		if prod.CategoryId == categoryId {
			result = append(result, prod)
		}
	}
	return result
}

func GetProduct(productId int) (*Product, error) {
	for _, p := range products {
		if p.Id == productId {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("Product not found with ID %v", productId)
}

var products []Product = []Product{Product{
	Name:             "Lemon Juice",
	DescriptionShort: "Made from fresh, organic California lemons.",
	DescriptionLong: `Made from premium, organic Meyer lemons. These fruit are left on the tree until they reach the peak of ripeness and then juiced within 8 hours of being picked.
			<br/>
			Lemonade made from our premium juice is sure to make your stand the most popular in the neighborhood.`,
	PricePerLiter:   1.09,
	PricePer10Liter: 1.04,
	Origin:          "California",
	IsOrganic:       true,
	ImageURL:        "lemon.png",
	Id:              1,
	CategoryId:      1,
}, Product{
	Name:             "Apple Juice",
	DescriptionShort: "The perfect blend of Washington apples.",
	DescriptionLong:  "The perfect blend of Washington apples.",
	PricePerLiter:    0.89,
	PricePer10Liter:  0.85,
	Origin:           "Ohio",
	IsOrganic:        true,
	ImageURL:         "apple.png",
	Id:               2,
	CategoryId:       1,
}, Product{
	Name:             "Watermelon Juice",
	DescriptionShort: "From sun-drenched fields in Florida.",
	DescriptionLong:  "From sun-drenched fields in Florida.",
	PricePerLiter:    3.99,
	PricePer10Liter:  3.84,
	Origin:           "Florida",
	IsOrganic:        true,
	ImageURL:         "watermelon.png",
	Id:               3,
	CategoryId:       1,
}, Product{
	Name:             "Kiwi Juice",
	DescriptionShort: "California sunshine and rain distilled into sweet goodness",
	DescriptionLong:  "California sunshine and rain distilled into sweet goodness",
	PricePerLiter:    5.99,
	PricePer10Liter:  5.54,
	Origin:           "California",
	IsOrganic:        false,
	ImageURL:         "kiwi.png",
	Id:               4,
	CategoryId:       1,
}, Product{
	Name:             "Mangosteen Juice",
	DescriptionShort: "Our latest taste sensation, imported directly from Hawaii",
	DescriptionLong:  "Our latest taste sensation, imported directly from Hawaii",
	PricePerLiter:    6.87,
	PricePer10Liter:  6.79,
	Origin:           "Hawaii",
	IsOrganic:        false,
	ImageURL:         "mangosteen.png",
	Id:               5,
	CategoryId:       1,
}, Product{
	Name:             "Orange Juice",
	DescriptionShort: "Fresh squeezed from Florida's best oranges.",
	DescriptionLong:  "Fresh squeezed from Florida's best oranges.",
	PricePerLiter:    1.67,
	PricePer10Liter:  1.63,
	Origin:           "Florida",
	IsOrganic:        false,
	ImageURL:         "orange.png",
	Id:               6,
	CategoryId:       1,
}, Product{
	Name:             "Pineapple Juice",
	DescriptionShort: "An exotic and refreshing offering. Straight from Hawaii.",
	DescriptionLong:  "An exotic and refreshing offering. Straight from Hawaii.",
	PricePerLiter:    2.48,
	PricePer10Liter:  2.27,
	Origin:           "Hawaii",
	IsOrganic:        false,
	ImageURL:         "pineapple.png",
	Id:               7,
	CategoryId:       1,
}, Product{
	Name:             "Strawberry Juice",
	DescriptionShort: "MThe perfect balance of sweet and tart.",
	DescriptionLong:  "The perfect balance of sweet and tart.",
	PricePerLiter:    4.36,
	PricePer10Liter:  4.27,
	Origin:           "California",
	IsOrganic:        false,
	ImageURL:         "strawberry.png",
	Id:               8,
	CategoryId:       1,
},
}
