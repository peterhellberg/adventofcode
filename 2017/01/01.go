package main

import "fmt"

var input string

func main() {
	fmt.Scan(&input)
	fmt.Printf("%q = %d and %d\n", input,
		inverseCaptchaNext(input),
		inverseCaptchaHalfwayAround(input),
	)
}

func inverseCaptchaNext(input string) (result int) {
	for i, r := range input {
		var ni int

		if i+1 < len(input) {
			ni = i + 1
		}

		if c := int(r) - 48; c == int(input[ni])-48 {
			result += c
		}
	}

	return result
}

func inverseCaptchaHalfwayAround(input string) (result int) {
	length, steps := len(input), len(input)/2

	for i, r := range input {
		c, ni := int(r)-48, i+steps

		if ni >= length {
			ni = i + steps - length
		}

		if c == int(input[ni])-48 {
			result += c
		}
	}

	return result
}
