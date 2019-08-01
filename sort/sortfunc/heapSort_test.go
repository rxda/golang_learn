package sortfunc

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_heapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"heap sort",
			args{
				[]int{4,5,6,2,1,3,9,8,7,0},
			},
			[]int{0,1,2,3,4,5,6,7,8,9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapSortP(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapSort() = %v, want %v", got, tt.want)
			}else{
				fmt.Println(got)
			}
		})
	}
}


func Test_heapify(t *testing.T) {
	type args struct {
		tree []int
		n    int
		i    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				//       4
				//    10   3
				//   5  1 2
				tree: []int{4,10,3,5,1,2},
				i:    0,
			},
		},
		{
			"2",
			args{
				//      3
				//    5   4
				// 10  7 9

				//res
				//       5
				//   10     4
				//  3  7  9
				tree: []int{3,5,4,10,7,9},
				n:    6,
				i:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildHeap(tt.args.tree)
			fmt.Println(tt.args.tree)
		})
	}
}