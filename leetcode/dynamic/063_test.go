package dynamic

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[][]int{{1, 0}}}, 0},
		{"case 2", args{[][]int{{1}}}, 0},
		{"case 3", args{[][]int{{0}}}, 1},
		{"case 4", args{[][]int{{1}, {0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
