package backdate

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"case 1", args{"aab"}, [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"case 2", args{"cdd"}, [][]string{{"c", "dd"}, {"c", "d", "d"}}},
		{"case 3", args{"a"}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
