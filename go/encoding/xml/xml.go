package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct{
	Name string
	Age uint8
}

func main()  {
	toXMl()
}

func toXMl(){
	p := Person{
		Name: "Han",
		Age:  18,
	}
	x, err := xml.Marshal(p)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(x))
}


