package arrays

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
		want []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{[]int{}, 3, []int{}}, 0},
		{"case 1", args{[]int{3, 2, 2, 3}, 3, []int{2, 2}}, 2},
		{"case 2", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); !reflect.DeepEqual(tt.args.nums[:got], tt.args.want) {
				t.Errorf("removeElement() = %v, want %v", tt.args.nums[:got], tt.args.want)
			}
		})
	}
}
