package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	input := scanInts(os.Stdin)
	fmt.Println("Steps:", one(input), "and", two(input))
}

func one(input []int) int {
	return steps(input, func(i int, m []int) int {
		j := i + m[i]

		m[i]++

		return j
	})
}

func two(input []int) int {
	return steps(input, func(i int, m []int) int {
		j := i + m[i]

		if m[i] > 2 {
			m[i]--
		} else {
			m[i]++
		}

		return j
	})
}

func steps(input []int, jump func(i int, m []int) int) (n int) {
	m := make([]int, len(input))

	copy(m, input)

	for i := 0; i < len(m); n++ {
		i = jump(i, m)
	}

	return n
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
