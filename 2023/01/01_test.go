package main

import (
	"strings"
	"testing"
)

func TestDay1(t *testing.T) {
	t.Run("One", func(t *testing.T) {
		input := scanInput(strings.NewReader(oneInput))

		if got, want := one(input), 142; got != want {
			t.Fatalf("one(input) = %d, want %d", got, want)
		}
	})

	t.Run("Two", func(t *testing.T) {
		input := scanInput(strings.NewReader(twoInput))

		if got, want := two(input), 281; got != want {
			t.Fatalf("two(input) = %d, want %d", got, want)
		}
	})
}

const oneInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const twoInput = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
