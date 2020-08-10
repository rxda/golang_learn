package main

import (
	"fmt"
	"golang_learn/unsafe_change_private field/in"
	"reflect"
	"unsafe"
)

// try to change private field of struct
func main() {
	book := in.NewBook()
	fmt.Println(book)
	// try to change price
	field := reflect.ValueOf(&book).Elem().FieldByName("price")
	SetUnexportedField(field, 9.9)
	fmt.Println(book)
}

func SetUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}