package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := scanInput(os.Stdin)

	fmt.Println("Part one:", one(input))
	fmt.Println("Part two:", two(input))
}

func one(input []int) int {
	if l := len(input); l > 0 {
		return input[l-1]
	}

	return 0
}

func two(input []int) int {
	var sum int

	if l := len(input); l > 2 {
		for _, e := range input[l-3:] {
			sum += e
		}
	}

	return sum
}

func scanInput(r io.Reader) []int {
	var elves []int
	var e int

	s := bufio.NewScanner(r)

	for s.Scan() {
		if s.Text() == "" {
			elves = append(elves, e)
			e = 0
			continue
		}

		if n, err := strconv.Atoi(s.Text()); err == nil {
			e += n
		}
	}

	if e > 0 {
		elves = append(elves, e)
	}

	sort.Ints(elves)

	return elves
}
