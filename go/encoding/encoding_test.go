package encoding

import "testing"

func Test_toJson(t *testing.T) {
	type args struct {
		d *Dog
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			"dd",
			args{d:&Dog{
				"happy",
				2,
				"black",
			}},
			"",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := toJson(tt.args.d)
			if got != tt.want {
				t.Errorf("toJson() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("toJson() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}