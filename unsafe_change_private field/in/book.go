package in

import "time"

type Book struct {
	Name   string
	Date   time.Time
	price  float64
	auther Author
}

type Author struct {
	Name string
	Age  uint8
}

func NewBook() Book {
	return Book{
		Name:   "golang110",
		Date:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		price:  99.5,
		auther: Author{
			Name: "Rx",
			Age:  18,
		},
	}
}
