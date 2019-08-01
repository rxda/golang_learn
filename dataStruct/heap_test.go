package dataStruct

import (
	"fmt"
	"ginwithmodule/sort/sortfunc"
	"testing"
)

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
				n:    6,
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
			Heapify(tt.args.tree,tt.args.n,tt.args.i)
			fmt.Println(tt.args.tree)
		})
	}
}

func Test_buildHead(t *testing.T) {
	type args struct {
		tree []int
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"2",
			args{
				//      3
				//    5   4
				// 10  7 9

				//res
				//       10
				//    7      9
				//  5  3  4
				tree: []int{3,5,4,10,7,9},
				n:    6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortfunc.Heapify(tt.args.tree, tt.args.n)
			fmt.Println(tt.args.tree)
		})
	}
}