package dynamic

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{1, 2, 3, 1}}, 4},
		{"case 2", args{[]int{2, 7, 9, 3, 1}}, 12},
		{"case 3", args{[]int{2, 1}}, 2},
		{"case 4", args{[]int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
