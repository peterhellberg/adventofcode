package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := scanInput(os.Stdin)

	fmt.Println("Part one:", one(input))
	fmt.Println("Part two:", two(input))
}

func one(games Games) int {
	return games.Sum(Round{12, 13, 14})
}

func two(games Games) int {
	return games.Min()
}

func scanInput(r io.Reader) (games Games) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		var g Game

		before, rest := cut(s.Text()[5:], ":")

		g.ID = atoi(before)

		for _, round := range strings.Split(rest, ";") {
			var r Round
			for _, p := range strings.Split(strings.TrimSpace(round), ", ") {
				switch s, c := cut(p, " "); c {
				case "red":
					r.R = atoi(s)
				case "green":
					r.G = atoi(s)
				case "blue":
					r.B = atoi(s)
				}
			}
			g.Rounds = append(g.Rounds, r)
		}
		games = append(games, g)
	}
	return games
}

type Round struct {
	R int
	G int
	B int
}

func (r Round) Power() int {
	return r.R * r.G * r.B
}

type Game struct {
	ID     int
	Rounds []Round
}

type Games []Game

func (games Games) Sum(max Round) (sum int) {
	for _, g := range games {
		ok := true

		for _, r := range g.Rounds {
			if r.R > max.R || r.G > max.G || r.B > max.B {
				ok = false
			}
		}

		if ok {
			sum += g.ID
		}
	}

	return sum
}

func (games Games) Min() (sum int) {
	for _, g := range games {
		var min Round

		for _, r := range g.Rounds {
			if r.R > min.R {
				min.R = r.R
			}
			if r.G > min.G {
				min.G = r.G
			}
			if r.B > min.B {
				min.B = r.B
			}
		}

		sum += min.Power()
	}

	return sum
}

func cut(s, sep string) (b, a string) {
	b, a, ok := strings.Cut(s, sep)
	if !ok {
		panic("could not cut")
	}
	return b, a
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
