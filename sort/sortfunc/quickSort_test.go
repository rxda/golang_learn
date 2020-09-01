package sortfunc

import (
	"fmt"
	"testing"
)

func Test_quickPart(t *testing.T) {
	tt := []int{'Q', 'H', 'C', 'Y', 'P', 'A', 'M', 'S', 'R', 'D', 'F', 'X'}
	for i := 0; i < len(tt)-1; i++ {
		for j := i + 1; j < len(tt); j++ {
			if tt[i] > tt[j] {
				tt[i], tt[j] = tt[j], tt[i]
			}
		}
		break
	}
	for _, v := range tt {
		fmt.Println(string(rune(v)))
	}
}

func quickPart2(arr []int, left, right int) []int {
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
	return arr
}
