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
	fmt.Println("Bottom tower:", one(input))
}

func one(towers []tower) string {
	var holding []tower

	held := map[string]bool{}

	for _, t := range towers {
		if t.Holding != nil {
			holding = append(holding, t)

			for _, h := range t.Holding {
				held[h] = true
			}
		}
	}

	var bottom string

	for _, h := range holding {
		if !held[h.Name] {
			bottom = h.Name
		}
	}

	return bottom
}

func scanInput(r io.Reader) (towers []tower) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		parts := strings.Split(replacer.Replace(s.Text()), " ")
		count := len(parts)

		if count == 0 {
			continue
		}

		t := tower{Name: parts[0]}

		if count > 1 {
			t.Weight, _ = strconv.Atoi(parts[1])
		}

		if count > 2 {
			t.Holding = parts[2:]
		}

		towers = append(towers, t)
	}

	return towers
}

type tower struct {
	Name    string   `json:"name"`
	Weight  int      `json:"weight"`
	Holding []string `json:"holding,omitempty"`
}

var replacer = strings.NewReplacer("(", "", ")", "", " ->", "", ",", "")
