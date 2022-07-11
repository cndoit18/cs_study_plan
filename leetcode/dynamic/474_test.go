package dynamic

import "testing"

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]string{"10", "0001", "111001", "1", "0"}, 5, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
