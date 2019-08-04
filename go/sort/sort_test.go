package sort

import "testing"

func Test_sortNumbers(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			"1",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortNumbers(); got != tt.want {
				t.Errorf("sortNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structSort(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			"1",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := structSort(); got != tt.want {
				t.Errorf("structSort() = %v, want %v", got, tt.want)
			}
		})
	}
}