package mysync

import "testing"

func Test_testOnce(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"1",
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testOnce(); got != tt.want {
				t.Errorf("testOnce() = %v, want %v", got, tt.want)
			}
		})
	}
}
