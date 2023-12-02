package main

import (
	"strings"
	"testing"
)

func TestDay2(t *testing.T) {
	t.Run("One", func(t *testing.T) {
		input := scanInput(strings.NewReader(oneInput))

		if got, want := one(input), 8; got != want {
			t.Fatalf("one(input) = %d, want %d", got, want)
		}
	})

	t.Run("Two", func(t *testing.T) {
		input := scanInput(strings.NewReader(oneInput))

		if got, want := two(input), 2286; got != want {
			t.Fatalf("two(input) = %d, want %d", got, want)
		}
	})
}

const oneInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
