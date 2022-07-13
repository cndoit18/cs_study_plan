package dynamic

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{1, 2, 5}, 11}, 3},
		{"case 2", args{[]int{2}, 3}, -1},
		{"case 3", args{[]int{2}, 4}, 2},
		{"case 4", args{[]int{186, 419, 83, 408}, 6249}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
