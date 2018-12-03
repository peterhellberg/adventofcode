package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	ints := scanInts(os.Stdin)

	fmt.Println("One:", one(ints))
	fmt.Println("Two:", two(ints))
}

func one(ints []int) (result int) {
	for i := range ints {
		result += ints[i]
	}

	return result
}

func two(ints []int) int {
	var result int

	f := map[int]int{}

	for {
		for i := range ints {
			result += ints[i]

			f[result]++

			if f[result] == 2 {
				return result
			}
		}
	}
}

func scanInts(r io.Reader) (ints []int) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		if i, err := strconv.Atoi(s.Text()); err == nil {
			ints = append(ints, i)
		}
	}

	return ints
}
