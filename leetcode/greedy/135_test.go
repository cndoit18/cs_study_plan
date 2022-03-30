package greedy

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{1, 0, 2}}, 5},
		{"case 2", args{[]int{1, 2, 2}}, 4},
		{"case 3", args{[]int{1, 3, 2, 2, 1}}, 7},
		{"case 4", args{[]int{1, 2, 87, 87, 87, 2, 1}}, 13},
		{"case 5", args{[]int{1, 3, 4, 5, 2}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
