package mysync

import "testing"

func Test_errorWg(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			"1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorWg(); got != tt.want {
				t.Errorf("errorWg() = %v, want %v", got, tt.want)
			}
		})
	}
}