package viewmodel

type Home struct {
	Title  string
	Active string
}

func NewHome() Home {
	result := Home{Title: "Lemonade Stand Supply",
		Active: "home",
	}
	return result
}
