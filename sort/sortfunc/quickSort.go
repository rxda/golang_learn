package sortfunc

func quick(arr []int) []int {
	quickPart(arr, 0, len(arr)-1)
	return arr
}

func quickPart(arr []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right

	for i < j {
		for arr[j] >= arr[left] && i < j {
			j--
		}

		for arr[i] <= arr[left] && i < j {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left], arr[i] = arr[i], arr[left]

	quickPart(arr, left, i-1)
	quickPart(arr, i+1, right)
}
