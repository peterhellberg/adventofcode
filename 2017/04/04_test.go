package main

import "testing"

func TestValid(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  bool
	}{
		{"", false},
		{"aa bb", true},
		{"aa bb aa", false},
		{"aa bb cc", true},
	} {
		if got := valid(tt.input); got != tt.want {
			t.Fatalf("valid(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
