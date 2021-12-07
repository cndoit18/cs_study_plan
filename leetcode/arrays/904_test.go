package arrays

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		fruits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{[]int{}}, 0},
		{"case 1", args{[]int{1, 2, 1}}, 3},
		{"case 2", args{[]int{0, 1, 2, 2}}, 3},
		{"case 3", args{[]int{1, 2, 3, 2, 2}}, 4},
		{"case 4", args{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.args.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
