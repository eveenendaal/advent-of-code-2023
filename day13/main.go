package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func mirror(s []string, equal func([]string, []string) bool) int {
	for i := 1; i < len(s); i++ {
		l := slices.Min([]int{i, len(s) - i})
		a, b := slices.Clone(s[i-l:i]), s[i:i+l]
		slices.Reverse(a)
		if equal(a, b) {
			return i
		}
	}
	return 0
}

func smudge(a, b []string) bool {
	diffs := 0
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				diffs++
			}
		}
	}
	return diffs == 1
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func transpose(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}

	result := make([]string, len(strings[0]))
	for _, s := range strings {
		for i, r := range s {
			result[i] = result[i] + string(r)
		}
	}

	return result
}

func handlePattern(rows []string, part2 bool) int {
	columns := transpose(rows)
	if part2 {
		return mirror(columns, smudge) + 100*mirror(rows, smudge)
	} else {
		return mirror(columns, equal) + 100*mirror(rows, equal)
	}
}

func Solve(filePath string, part2 bool) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		} else {
			total += handlePattern(lines, part2)
			lines = []string{}
		}
	}
	total += handlePattern(lines, part2)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	return total
}

func main() {
	fmt.Println("Advent of Code 2017 - Day 13")
	// fmt.Println("Part 1:", Solve("data.txt", false))
	fmt.Println("Part 2:", Solve("data.txt", true))
}
