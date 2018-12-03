package main

import (
	"strings"
	"testing"
)

func TestDay1(t *testing.T) {
	t.Run("One", func(t *testing.T) {
		input := scanInts(strings.NewReader(oneInput))

		if got, want := one(input), -24; got != want {
			t.Fatalf("one(input) = %d, want %d", got, want)
		}
	})

	t.Run("Two", func(t *testing.T) {
		input := scanInts(strings.NewReader(twoInput))

		if got, want := two(input), 14; got != want {
			t.Fatalf("two(input) = %d, want %d", got, want)
		}
	})
}

var oneInput = `+4
-18
+8
+16
-13
-10
-11
`

var twoInput = `+7
+7
-2
-7
-4
`
