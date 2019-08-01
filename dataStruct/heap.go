package dataStruct

type heap []int
// left = 2*i+1
//right = 2*i+2
//parent = (i-1)/2


func Heapify(tree []int, n, i int) {
	//i=2,c1=5,c2=6
	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}
	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}

	if max != i{
		tree[max], tree[i] = tree[i], tree[max]
		Heapify(tree, n , max)
	}
}

func BuildHeap(tree []int, n int){
	lastNode := n-1
	parent := (lastNode-1) /2
	for i:= parent;i>=0;i--{
		Heapify(tree,n,i)
	}
}
