package pointer

import (
	"fmt"
	"strconv"
)
type SA []*A

type A struct{
	Name string
	Age uint8
}

func tttt()  {
	s := make([]A, 4)
	for k := range s{
		s[k] = A{
			Name: strconv.Itoa(k),
			Age:  uint8(k),
		}
	}
	for k,v := range s{

	}
}

func (a *SA) t1() {
	([]*A(*a))[0].Age ++
}

