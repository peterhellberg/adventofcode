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
	sp := makeSpreadsheet(os.Stdin)
	fmt.Println(sp.checksum(), "and", sp.evenlyDivisibleChecksum())
}

type spreadsheet struct {
	data [][]int
}

func makeSpreadsheet(r io.Reader) spreadsheet {
	scanner := bufio.NewScanner(r)

	sp := spreadsheet{}

	for scanner.Scan() {
		var d []int

		for _, s := range strings.Split(scanner.Text(), "\t") {
			if i, err := strconv.Atoi(s); err == nil {
				d = append(d, i)
			}
		}

		sp.data = append(sp.data, d)
	}

	return sp
}

func (s spreadsheet) checksum() (c int) {
	for _, row := range s.data {
		var min, max int

		for i, n := range row {
			switch {
			case i == 0:
				min = n
				max = n
			case n < min:
				min = n
			case n > max:
				max = n
			}
		}

		c += max - min
	}

	return c
}

func (s spreadsheet) evenlyDivisibleChecksum() (c int) {
	for _, row := range s.data {
		for i, n := range row {
			for j, o := range row {
				if i != j && n%o == 0 {
					c += n / o
				}
			}
		}
	}

	return c
}
