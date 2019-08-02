package dfsbfs

import (
	"reflect"
	"testing"
)

func Test_bfs(t *testing.T) {
	type args struct {
		g     map[string][]string
		start string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"begin from A",
			args{
				g:map[string][]string{
					"A": {"B", "C"},
					"B": {"A", "C", "D"},
					"C": {"A", "B", "D", "E"},
					"D": {"B", "C", "E", "F"},
					"E": {"C", "D"},
					"F": {"D"},
				},
				start :"A",
			},
			[]string{"A","B","C","D","E","F"},
		},
		{
			"begin from E",
			args{
				g:map[string][]string{
					"A": {"B", "C"},
					"B": {"A", "C", "D"},
					"C": {"A", "B", "D", "E"},
					"D": {"B", "C", "E", "F"},
					"E": {"C", "D"},
					"F": {"D"},
				},
				start :"E",
			},
			[]string{"E","C","D","A","B","F"},
		},
		{
			"begin from D",
			args{
				g:map[string][]string{
					"A": {"B", "C"},
					"B": {"A", "C", "D"},
					"C": {"A", "B", "D", "E"},
					"D": {"B", "C", "E", "F"},
					"E": {"C", "D"},
					"F": {"D"},
				},
				start :"D",
			},
			[]string{"D","B","C","E","F","A"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bfs(tt.args.g, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
