package hashtable

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{[]int{}, []int{}}, []int{}},
		{"case 1", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{9, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(sort.IntSlice(got), sort.IntSlice(tt.want)) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
