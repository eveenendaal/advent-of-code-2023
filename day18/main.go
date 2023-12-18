package main

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
)

func Solve(filepath string) int {
	total := 0

	lines := aoc.ReadFileToLines(filepath)
	for _, line := range lines {
		fmt.Printf("Line: %s\n", line)
	}

	return total

}
