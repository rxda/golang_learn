package mysync

import "testing"

//fatal err and can't recover
func TestErrCoroutineMap(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			"error",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrCoroutineMap(); got != tt.want {
				t.Errorf("ErrCoroutineMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCorrectCoroutineMap(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			"success",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CorrectCoroutineMap(); got != tt.want {
				t.Errorf("CorrectCoroutineMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOtherCoroutineMapFunc(t *testing.T) {
	tests := []struct {
		name       string
		wantDelete bool
		wantV100   int
		wantV1001  int
	}{
		{
			"1",
			true,
			11,
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDelete, gotV100, gotV1001 := OtherCoroutineMapFunc()
			if gotDelete != tt.wantDelete {
				t.Errorf("OtherCoroutineMapFunc() gotDelete = %v, want %v", gotDelete, tt.wantDelete)
			}
			if gotV100 != tt.wantV100 {
				t.Errorf("OtherCoroutineMapFunc() gotV100 = %v, want %v", gotV100, tt.wantV100)
			}
			if gotV1001 != tt.wantV1001 {
				t.Errorf("OtherCoroutineMapFunc() gotV1001 = %v, want %v", gotV1001, tt.wantV1001)
			}
		})
	}
}
