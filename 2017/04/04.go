package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if valid(scanner.Text()) {
			count++
		}
	}

	fmt.Println("Valid:", count)
}

func valid(s string) bool {
	if s == "" {
		return false
	}

	seen := map[string]int{}

	for _, word := range strings.Split(s, " ") {
		seen[word]++

		if seen[word] > 1 {
			return false
		}
	}

	return true
}
