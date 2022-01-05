package string

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{"aabaabaafa", "aabaaf"}, 3},
		{"case 2", args{"aaaaaaa", "aa"}, 0},
		{"case 3", args{"aabaaabaaac", "aabaaac"}, 4},
		{"case 4", args{"baabbbbababbbabab", "abbab"}, -1},
		{"case 5", args{"ababcaababcaabc", "ababcaabc"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
