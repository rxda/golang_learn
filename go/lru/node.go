package lru

type Node struct {
	// key
	Key interface{}
	// value
	Value interface{}
	// pre node
	pre *Node
	// new node
	next *Node
}

func NewNode(key, value interface{}) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}
