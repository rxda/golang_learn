package mycontainer

import "testing"

func Test_ringLearn(t *testing.T) {
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
			if got := ringLearn(); got != tt.want {
				t.Errorf("ringLearn() = %v, want %v", got, tt.want)
			}
		})
	}
}