package main

import (
	"strings"
	"testing"
)

func TestCorruptionChecksum(t *testing.T) {
	var input = strings.NewReader("5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8\n")

	if got, want := makeSpreadsheet(input).checksum(), 18; got != want {
		t.Fatalf("makeSpreadsheet(input).checksum() = %d, want %d", got, want)
	}
}

func TestEvenlyDivisibleChecksum(t *testing.T) {
	var input = strings.NewReader("5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5\n")

	if got, want := makeSpreadsheet(input).evenlyDivisibleChecksum(), 9; got != want {
		t.Fatalf("makeSpreadsheet(input).evenlyDivisibleChecksum() = %d, want %d", got, want)
	}
}
