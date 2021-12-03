package arrays

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
		want []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{[]int{}, []int{}}, 0},
		{"case 1", args{[]int{0, 0}, []int{0}}, 1},
		{"case 2", args{[]int{1, 1, 2}, []int{1, 2}}, 2},
		{"case 3", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); !reflect.DeepEqual(tt.args.nums[:got], tt.args.want) || tt.want != got {
				t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums[:got], tt.args.want)
			}
		})
	}
}
