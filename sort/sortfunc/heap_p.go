package sortfunc

func heapify(arr []int, i int) {
	n := len(arr)
	c1, c2 := 2*i+1, 2*i+2
	//比较得到一个父节点i和两个子节点c1,c2的最大值
	max := i
	if c1 < n && arr[c1] > arr[max] {
		max = c1
	}
	if c2 < n && arr[c2] > arr[max] {
		max = c2
	}

	if max != i {
		//最大值和父节点i交换
		arr[max], arr[i] = arr[i], arr[max]
		heapify(arr, max)
	}

}

func buildHeap(arr []int) {
	n := len(arr)
	//从最后一个节点的父节点开始堆化
	for i := (n - 2) / 2; i >= 0; i-- {
		heapify(arr, i)
	}
}

func heapSortP(arr []int) []int {
	//构建堆
	buildHeap(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		//root节点与最后一个节点交换，最后一个节点为当前最大值
		arr[i], arr[0] = arr[0], arr[i]
		//堆化root节点
		heapify(arr[:i], 0)
	}
	return arr
}
