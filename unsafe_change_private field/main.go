package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

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

// try to change private field of struct
func main() {
	book := NewBook()
	fmt.Println(book)
	// try to change price
	field := reflect.ValueOf(&book).Elem().FieldByName("price")
	SetUnexportedField(field, 9.9)
	fmt.Println(book)
}

func SetUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}