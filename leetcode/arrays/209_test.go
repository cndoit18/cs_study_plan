package arrays

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{3, []int{}}, 0},
		{"case 1", args{7, []int{2, 3, 1, 2, 4, 3}}, 2},
		{"case 2", args{4, []int{1, 4, 4}}, 1},
		{"case 3", args{11, []int{1, 2, 3, 4, 5}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
