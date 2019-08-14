package tree

//二叉查找树
type T int
type Node struct {
	Val   T
	Left  *Node
	Right *Node
}

type BST struct {
	root *Node
}

func BSTMaker() *BST {
	return &BST{}
}

func (b *Node) Insert(val T) {
	if val > b.Val {
		if b.Right == nil {
			b.Right = &Node{Val: val}
		} else {
			b.Right.Insert(val)
		}
	} else if val < b.Val {
		if b.Left == nil {
			b.Left = &Node{Val: val}
		}
	}
}

func (t *BST) Insert(val T){

}