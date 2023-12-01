package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := scanInput(os.Stdin)

	fmt.Println("Part one:", one(input))
	fmt.Println("Part two:", two(input))
}

func one(lines []string) (result int) {
	for _, line := range lines {
		result += value(line)
	}
	return result
}

func two(lines []string) (result int) {
	prefix := func(s, w, d string) string {
		if strings.HasPrefix(s, w) || strings.HasPrefix(s, d) {
			return d
		}
		return ""
	}

	for _, line := range lines {
		var d string
		for i := range line {
			r := string(line[i:])
			d += prefix(r, "one", "1")
			d += prefix(r, "two", "2")
			d += prefix(r, "three", "3")
			d += prefix(r, "four", "4")
			d += prefix(r, "five", "5")
			d += prefix(r, "six", "6")
			d += prefix(r, "seven", "7")
			d += prefix(r, "eight", "8")
			d += prefix(r, "nine", "9")
		}
		result += value(d)
	}
	return result
}

func value(s string) int {
	var first, last int
	for _, r := range s {
		if r > '0' && r <= '9' {
			if first == 0 {
				first = int(r - '0')
			}
			last = int(r - '0')
		}
	}
	return (first * 10) + last
}

func scanInput(r io.Reader) (lines []string) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}
