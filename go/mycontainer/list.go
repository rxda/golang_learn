package mycontainer

import (
	"container/list"
	"fmt"
)

func listLearn(){
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4) //4
	e1 := l.PushFront(1) //1 4
	l.InsertBefore(3, e4)//1 3 4
	l.InsertAfter(2, e1)// 1 2 3 4

	// Iterate through list and print its contents.
	//loop e = e.Next()
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}