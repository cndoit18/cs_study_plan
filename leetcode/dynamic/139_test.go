package dynamic

import "testing"

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 1", args{"leetcode", []string{"leet", "code"}}, true},
		{"case 2", args{"aaaaaaa", []string{"aaaa", "aaa"}}, true},
		{"case 3", args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
