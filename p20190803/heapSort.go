package p20190803

func heapify(arr []int, i int) {
	//两个子节点
	c1 := 2*i + 1
	c2 := c1 + 1
	//比较三个谁大
	max := i
	//--error : 注意判断根节点是否存在
	// if arr[c1] > arr[max] {
	// 	max = c1
	// }
	// if arr[c2] > arr[max] {
	// 	max = c2
	// }
	if c1 < len(arr) && arr[c1] > arr[max] {
		max = c1
	}
	if c2 < len(arr) && arr[c2] > arr[max] {
		max = c2
	}

	//如果最大的不是当前的，交换当前的和最大的的值,继续堆化max
	if max != i {
		arr[i], arr[max] = arr[max], arr[i]
		heapify(arr, max)
	}
}

func buildHeap(arr []int) {
	//找到最后一个节点的父节点
	last := len(arr) - 1
	parent := (last - 1) / 2
	//依次堆化
	for i := parent; i >= 0; i-- {
		heapify(arr, i)
	}
}

func heapSort(arr []int) []int {
	//构造堆
	buildHeap(arr)
	last := len(arr) - 1
	for i := last; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i)
	}
	return arr
}
