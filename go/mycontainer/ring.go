package mycontainer

import (
	"container/ring"
	"fmt"
)

func ringLearn() bool{
	// Create a new ring of size 5
	r := ring.New(6)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	fmt.Println(r.Value)
	// Iterate through the ring and print its contents
	//r.Do(func(p interface{}) {
	//	fmt.Println(p.(int))
	//})
	return true
}