package viewmodel

type SignUp struct {
	Title     string
	Active    string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func NewSignUp() SignUp {
	result := SignUp{
		Active: "home",
		Title:  "Lemonade Stand Supply",
	}
	return result
}
