package mysync

import "testing"

func Test_poolTest(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"1",
			"han,18",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := poolTest(); got != tt.want {
				t.Errorf("poolTest() = %v, want %v", got, tt.want)
			}
		})
	}
}