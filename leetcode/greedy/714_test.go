package greedy

import "testing"

func Test_maxProfit_714(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{1, 2, 3, 4, 5}, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit_714(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit_714() = %v, want %v", got, tt.want)
			}
		})
	}
}
