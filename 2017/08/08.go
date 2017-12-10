package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	input := scanInput(os.Stdin)
	fmt.Println("Largest value:", one(input), "and highest:", two(input))
}

func one(instructions []instruction) (largest int) {
	for _, v := range run(instructions) {
		if v > largest {
			largest = v
		}
	}

	return largest
}

func two(instructions []instruction) (highest int) {
	run(instructions, func(r registers, i instruction) {
		if r[i.reg] > highest {
			highest = r[i.reg]
		}
	})

	return highest
}

func run(instructions []instruction, funcs ...func(r registers, i instruction)) registers {
	r := registers{}

	for _, i := range instructions {
		switch {
		case
			i.ifcmp == "==" && r[i.ifreg] == i.ifval,
			i.ifcmp == "!=" && r[i.ifreg] != i.ifval,
			i.ifcmp == "<=" && r[i.ifreg] <= i.ifval,
			i.ifcmp == ">=" && r[i.ifreg] >= i.ifval,
			i.ifcmp == "<" && r[i.ifreg] < i.ifval,
			i.ifcmp == ">" && r[i.ifreg] > i.ifval:
			switch i.dir {
			case "inc":
				r[i.reg] += i.val
			case "dec":
				r[i.reg] -= i.val
			}

			for _, f := range funcs {
				f(r, i)
			}
		}
	}

	return r
}

func scanInput(r io.Reader) (instructions []instruction) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		i := instruction{}

		fmt.Sscanf(s.Text(), "%s %s %d if %s %s %d",
			&i.reg, &i.dir, &i.val, &i.ifreg, &i.ifcmp, &i.ifval)

		instructions = append(instructions, i)
	}

	return instructions
}

type registers map[string]int

type instruction struct {
	reg   string
	dir   string
	val   int
	ifreg string
	ifcmp string
	ifval int
}
