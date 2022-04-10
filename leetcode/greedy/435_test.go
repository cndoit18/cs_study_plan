package greedy

import "testing"

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}}, 1},
		{"case 2", args{[][]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}}}, 3},
		{"case 3", args{[][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
