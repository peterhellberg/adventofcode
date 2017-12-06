package main

import "testing"

func TestSteps(t *testing.T) {
	for _, tt := range []struct {
		input int
		want  int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 1},
		{7, 2},
		{8, 1},
		{9, 2},
		{10, 3},
		{11, 2},
		{12, 3},
		{13, 4},
		{14, 3},
		{15, 2},
		{16, 3},
		{17, 4},
		{18, 3},
		{19, 2},
		{20, 3},
		{21, 4},
		{22, 3},
		{23, 2},
		{24, 3},
		{25, 4},
		{26, 5},
	} {
		if got := steps(tt.input); got != tt.want {
			t.Fatalf("steps(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
