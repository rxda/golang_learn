package sortfunc

import "ginwithmodule/dataStruct"

func heapSort(arr []int) []int{
	n := len(arr)
	dataStruct.BuildHeap(arr, n)
	for i:=n-1;i>=0;i--{
		arr[i], arr[0] = arr[0], arr[i]
		dataStruct.Heapify(arr, i, 0)
	}
	return arr
}

