package base

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadFileToLines reads a file into a slice of strings
func ReadFileToLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// ReadFileToCharacters reads a file into a 2D array of runes
func ReadFileToCharacters(filePath string) [][]rune {
	lines := ReadFileToLines(filePath)
	characters := make([][]rune, len(lines))
	for y, line := range lines {
		for x, char := range line {
			if characters[y] == nil {
				characters[y] = make([]rune, len(line))
			}
			characters[y][x] = char
		}
	}
	return characters
}

// ReadFileToInt reads a file into a slice of ints
func ReadFileToInt(filePath string) [][]int {
	lines := ReadFileToLines(filePath)
	results := make([][]int, len(lines))
	for y, line := range lines {
		for x, char := range line {
			if results[y] == nil {
				results[y] = make([]int, len(line))
			}
			results[y][x], _ = strconv.Atoi(string(char))
		}
	}
	return results
}

type Position struct {
	Column int
	Row    int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// IntAbs Absolute value of an integer
func IntAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
