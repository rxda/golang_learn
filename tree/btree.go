package tree

import "fmt"

type BinaryTree struct {
	Val int
	Left *BinaryTree
	Right *BinaryTree

}

// LayerTravers 二叉树按层遍历
func LayerTravers(root *BinaryTree)  {
	var que = make([]*BinaryTree,0,4)
	que = append(que, root)

	for len(que) > 0{
		tmp := que[0]
		if tmp.Left != nil{
			que = append(que, tmp.Left)
		}
		if tmp.Right != nil{
			que = append(que, tmp.Right)
		}
		fmt.Println(que[0].Val)
		que = que[1:]
	}
}
//       6
//    3     1
//  2   7
// 8

func BuildTree()  BinaryTree{
	a := BinaryTree{
		Val:   6,
		Left:  &BinaryTree{
			Val:   3,
			Left:  &BinaryTree{
				Val:   2,
				Left:  &BinaryTree{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &BinaryTree{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryTree{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	return a
}

//normal tree



