package main

import (
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	input := scanInput(strings.NewReader(oneList))

	if got, want := one(input), "tknk"; got != want {
		t.Fatalf("one(input) = %q, want %q", got, want)
	}
}

const oneList = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)
`
