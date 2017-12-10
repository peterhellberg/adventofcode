package main

import (
	"strings"
	"testing"
)

func TestDay8(t *testing.T) {
	input := scanInput(strings.NewReader(inputString))

	t.Run("One", func(t *testing.T) {
		if got, want := one(input), 1; got != want {
			t.Fatalf("one(input) = %d, want %d", got, want)
		}
	})

	t.Run("Two", func(t *testing.T) {
		if got, want := two(input), 10; got != want {
			t.Fatalf("two(input) = %d, want %d", got, want)
		}
	})
}

const inputString = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
`
