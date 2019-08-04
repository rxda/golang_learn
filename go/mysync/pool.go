package mysync

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  uint8
}

func (p *Person) String() string {
	return fmt.Sprintf("%s,%d", p.Name, p.Age)
}

var personPool = sync.Pool{
	New: func() interface{} {
		return new(Person)
	},
}

func poolTest() string {
	person := personPool.Get().(*Person)
	person.Name = "han"
	person.Age = 18
	//s := fmt.Sprint(person)
	person = &Person{}
	personPool.Put(person)
	pp := personPool.Get().(*Person)
	s := fmt.Sprint(pp)
	return s

}
