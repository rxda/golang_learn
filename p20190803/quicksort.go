package p20190803

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right

	for i < j {
		for i < j && arr[j] >= arr[left] {
			j--
		}
		for i < j && arr[i] <= arr[left] {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[left], arr[i] = arr[i], arr[left]

	qs(arr, left, i-1)
	qs(arr, i+1, right)

}
