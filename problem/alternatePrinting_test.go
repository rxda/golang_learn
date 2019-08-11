package problem

import "testing"


func TestAlternatePrint1(t *testing.T) {
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
			if got := AlternatePrint(); got != tt.want {
				t.Errorf("AlternatePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}