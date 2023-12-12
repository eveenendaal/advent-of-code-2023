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

func dp(i, j int, record string, group []int, cache [][]int) int {

	if i >= len(record) {
		if j < len(group) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	fmt.Printf("i: %d, j: %d, record: %s, group: %v\n", i, j, record, group)

	res := 0
	if record[i] == '.' {
		// if character is a dot, we can skip it
		res = dp(i+1, j, record, group, cache)
	} else {
		// if character is a ?, we can skip it
		if record[i] == '?' {
			res += dp(i+1, j, record, group, cache)
		}
		// if we still have groups left
		if j < len(group) {
			count := 0
			// count the number of consecutive characters from current i
			for k := i; k < len(record); k++ {
				// stop if the count is greater then the group, you hit a ., or you hit the count with ?s
				if count > group[j] || record[k] == '.' || count == group[j] && record[k] == '?' {
					break
				}
				count += 1
			}

			// If the count matches the group, we can continue
			if count == group[j] {
				// if you haven't hit the end of the record, and the next character is not a #
				if i+count < len(record) && record[i+count] != '#' {
					res += dp(i+count+1, j+1, record, group, cache)
				} else {
					// if you have hit the end of the record, or the next character is a #
					res += dp(i+count, j+1, record, group, cache)
				}
			}
		}
	}

	cache[i][j] = res
	return res
}

func findSolutions(input string, counts []int, remaining int, unknowns int, total int) int {
	if unknowns == 0 {
		if remaining == 0 && unknowns == 0 {
			// fmt.Printf("Checking: %s\n", input)
			if validSolution(input, counts) {
				total += 1
			}
		}
		return total
	}
	// find next unknonwn
	for i := 0; i < len(input); i++ {
		if input[i] == '?' {
			// Replace chracter with # and .
			option1 := input[:i] + "#" + input[i+1:]
			option2 := input[:i] + "." + input[i+1:]

			if remaining > 0 {
				total = findSolutions(option1, counts, remaining-1, unknowns-1, total)
			}
			total = findSolutions(option2, counts, remaining, unknowns-1, total)
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
	// solutions := findSolutions(input, counts, remaining, unknowns, 0)
	var cache [][]int
	for i := 0; i < len(input); i++ {
		cache = append(cache, make([]int, len(counts)+1))
		for j := 0; j < len(counts)+1; j++ {
			cache[i][j] = -1
		}
	}
	solutions := dp(0, 0, input, counts, cache)

	fmt.Printf("Solutions: %d\n", solutions)
	if solutions == 0 {
		panic("No solutions found")
	}
	return solutions
}

func Solve(filePath string, part2 bool) int {
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
		springInput := input[0]

		if part2 {
			tempInput := []string{}
			tempCounts := []int{}
			for i := 0; i < 5; i++ {
				tempInput = append(tempInput, springInput)
				tempCounts = append(tempCounts, counts...)
			}
			springInput = strings.Join(tempInput, "?")
			counts = tempCounts
		}
		total += parseLine(springInput, counts)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return total
}

func main() {
	fmt.Println("Part 1:", Solve("data.txt", false))
	fmt.Println("Part 2:", Solve("data.txt", true))
}
