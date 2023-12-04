package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Part1(filepath string) {
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	total := 0

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line here
		value := processLine(line)
		total += value
	}

	fmt.Println("Total:", total)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func processLine(line string) int {
	// split line on ":"
	parts := strings.Split(line, ":")

	// remove all non digits from first part to get the game
	game := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, parts[0])

	// split the second part by the "|"
	input := strings.Split(parts[1], "|")

	// replace all the double spaces with single spaces
	input[0] = strings.Replace(input[0], "  ", " ", -1)
	input[1] = strings.Replace(input[1], "  ", " ", -1)

	// trim the first part of input and split on spaces
	winningNumbersStrings := strings.Split(strings.TrimSpace(input[0]), " ")
	foundNumbersString := strings.Split(strings.TrimSpace(input[1]), " ")

	// convert winning numbers to int
	winningNumbers := make([]int, len(winningNumbersStrings))
	for i, v := range winningNumbersStrings {
		winningNumbers[i], _ = strconv.Atoi(v)
	}
	// convert found numbers to int
	foundNumbers := make([]int, len(foundNumbersString))
	for i, v := range foundNumbersString {
		foundNumbers[i], _ = strconv.Atoi(v)
	}

	// How many numbers from the found numbers are in the winning numbers
	count := 0
	for _, v := range foundNumbers {
		for _, w := range winningNumbers {
			if v == w {
				count++
			}
		}
	}

	// Print variables
	fmt.Printf("Game: %s, Matches: %d, Winning: %v, Found: %v\n", game, count, winningNumbers, foundNumbers)

	if count == 0 {
		return 0
	} else {
		return int(math.Pow(2, float64(count-1)))
	}

}

func main() {
	Part1("data.txt")

}
