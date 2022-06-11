package dynamic

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{1, 1, 1, 1, 1}, 3}, 5},
		{"case 2", args{[]int{1}, 1}, 1},
		{"case 3", args{[]int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
