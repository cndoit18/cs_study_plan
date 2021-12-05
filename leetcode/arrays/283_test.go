package arrays

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{[]int{}}, []int{}},
		{"case 1", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if moveZeroes(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("searchRange() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
