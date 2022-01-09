package stack

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case 1", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []int{3, 3, 5, 5, 6, 7}},
		{"case 1", args{[]int{1, 3, 1, 2, 0, 5}, 3}, []int{3, 3, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
