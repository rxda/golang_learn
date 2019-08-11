package problem

import "testing"

func Test_builderTest(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"1",
			1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builderTest(); got != tt.want {
				t.Errorf("builderTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_safeBuilderTest(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"1",
			1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safeBuilderTest(); got != tt.want {
				t.Errorf("safeBuilderTest() = %v, want %v", got, tt.want)
			}
		})
	}
}