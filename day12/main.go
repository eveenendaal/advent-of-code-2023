package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func charCount(input string, char byte) int {
	count := 0
	for i := 0; i < len(input); i++ {
		if input[i] == char {
			count++
		}
	}
	return count
}

var spacesRegex = regexp.MustCompile(`\.+`)

func validSolution(input string, counts []int) bool {
	if charCount(input, '?') > 0 {
		return false
	}
	// trim '.' from input
	input = strings.Trim(input, ".")
	parts := spacesRegex.Split(input, -1)

	if len(parts) != len(counts) {
		return false
	}

	for i, v := range parts {
		if len(v) != counts[i] {
			return false
		}
	}
	return true
}

func findSolutions(input string, counts []int, remaining int, total int) int {
	if remaining == 0 {
		if validSolution(input, counts) {
			total += 1
		}
	}
	// find next unknonwn
	for i := 0; i < len(input); i++ {
		if input[i] == '?' {
			// Replace chracter with # and .
			option1 := input[:i] + "#" + input[i+1:]
			option2 := input[:i] + "." + input[i+1:]

			total = findSolutions(option1, counts, remaining-1, total)
			total = findSolutions(option2, counts, remaining, total)
			break
		}
	}
	return total
}

func parseLine(input string, counts []int) int {
	unknowns := charCount(input, '?')
	total := 0
	for _, count := range counts {
		total += count
	}
	on := charCount(input, '#')
	remaining := total - on

	fmt.Printf("Input: %s, Unknowns: %d, Remaining: %d, Counts: %v\n", input, unknowns, remaining, counts)
	solutions := findSolutions(input, counts, remaining, 0)
	fmt.Printf("Solutions: %d\n", solutions)
	if solutions == 0 {
		panic("No solutions found")
	}
	return solutions
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// split on whitespace
		input := strings.Split(line, " ")
		countString := strings.Split(input[1], ",")
		counts := make([]int, len(countString))
		for i, v := range countString {
			counts[i], _ = strconv.Atoi(v)
		}
		total += parseLine(input[0], counts)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return total
}

func main() {
	fmt.Println("Part 1:", Part1("data.txt"))
}
