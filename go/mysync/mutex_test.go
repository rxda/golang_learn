package mysync

import "testing"


//go test -run Test_errorIncrement go/mysync/mutex_test.go go/mysync/mutex.go
func Test_errorIncrement(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"error increment",
			1000,
		},
		{
			"error increment",
			1000,
		},
		{
			"error increment",
			1000,
		},
		{
			"error increment",
			1000,
		},
		{
			"error increment",
			1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorIncrement(); got != tt.want {
				t.Errorf("errorIncrement() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go test -run TestCorrectIncrement go/mysync/mutex_test.go go/mysync/mutex.go
func TestCorrectIncrement(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"correct increment",
			1000,
		},
		{
			"correct increment",
			1000,
		},
		{
			"correct increment",
			1000,
		},
		{
			"correct increment",
			1000,
		},
		{
			"correct increment",
			1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CorrectIncrement(); got != tt.want {
				t.Errorf("CorrectIncrement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testMutex(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"1",
			1000,
		},
		{
			"2",
			1000,
		},
		{
			"3",
			1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testMutex(); got != tt.want {
				t.Errorf("testMutex() = %v, want %v", got, tt.want)
			}
		})
	}
}