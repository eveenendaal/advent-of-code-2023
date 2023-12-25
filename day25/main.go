package day25

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
)

func Part1(filePath string) int {
	lines := aoc.ReadFileToLines(filePath)
	for _, line := range lines {
		fmt.Printf("line: %s\n", line)
	}
	return 0
}
