package main

import (
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	var ints, want = scanInts(strings.NewReader("0\n3\n0\n1\n-3\n")), 5

	if got := one(ints); got != want {
		t.Fatalf("one(%v) = %d, want %d", ints, got, want)
	}
}

func TestTwo(t *testing.T) {
	var ints, want = scanInts(strings.NewReader("0\n3\n0\n1\n-3\n")), 10

	if got := two(ints); got != want {
		t.Fatalf("two(%v) = %d, want %d", ints, got, want)
	}
}
