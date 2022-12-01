package main

import (
	"strings"
	"testing"
)

func TestDay1(t *testing.T) {
	input := scanInput(strings.NewReader(inputString))

	t.Run("One", func(t *testing.T) {
		if got, want := one(input), 24000; got != want {
			t.Fatalf("one(input) = %d, want %d", got, want)
		}
	})

	t.Run("Two", func(t *testing.T) {
		if got, want := two(input), 45000; got != want {
			t.Fatalf("two(input) = %d, want %d", got, want)
		}
	})
}

const inputString = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
