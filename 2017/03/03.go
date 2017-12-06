package main

import (
	"fmt"
	"math"
)

func main() {
	var input int

	fmt.Scan(&input)
	fmt.Println(steps(input))
}

func steps(input int) int {
	n, f := int(math.Ceil(math.Sqrt(float64(input)))), o

	if input%2 == 0 {
		f = e
	}

	var x0, x1, y0, y1 int

	for c := 0; c < n; c++ {
		for r := 0; r < n; r++ {
			switch f(n, r, c) + 1 {
			case 1:
				x1, y1 = c, r
			case input:
				x0, y0 = c, r
			}
		}
	}

	return abs(x0-x1) + abs(y0-y1)
}

func o(n, r, c int) int {
	if r == 0 {
		return n*n - 1 - c
	} else if c == n-1 {
		return n*n - 1 - c - r
	}

	return e(n-1, r-1, c)
}

func e(n, r, c int) int {
	if c == 0 {
		return (n-1)*(n-1) + r
	} else if r == n-1 {
		return (n-1)*(n-1) + r + c
	}

	return o(n-1, r, c-1)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
