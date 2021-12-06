package arrays

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{"", ""}, true},
		{"case 1", args{"ab#c", "ad#c"}, true},
		{"case 2", args{"ab##", "c#d#"}, true},
		{"case 3", args{"a##c", "#a#c"}, true},
		{"case 4", args{"a#c", "b"}, false},
		{"case 5", args{"xywrrmp", "xywrrmu#p"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
