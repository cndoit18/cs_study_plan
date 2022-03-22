package greedy

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 1", args{[]int{2, 3, 1, 1, 4}}, true},
		{"case 2", args{[]int{3, 2, 1, 0, 4}}, false},
		{"case 3", args{[]int{2, 0, 0}}, true},
		{"case 4", args{[]int{0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
