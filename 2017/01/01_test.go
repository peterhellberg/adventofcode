package main

import "testing"

func TestInverseCaptchaNext(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	} {
		if got := inverseCaptchaNext(tt.input); got != tt.want {
			t.Fatalf("inverseCaptchaNext(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestInverseCaptchaHalfwayAround(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	} {
		if got := inverseCaptchaHalfwayAround(tt.input); got != tt.want {
			t.Fatalf("inverseCaptchaHalfwayAround(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
