package tree

import (
	"fmt"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	b := new(Node)
	b.Insert(3)
	b.Insert(4)
	b.Insert(5)
	b.Insert(7)
	b.Insert(9)
	b.Insert(1)
	b.Insert(0)
	fmt.Println(b)
}