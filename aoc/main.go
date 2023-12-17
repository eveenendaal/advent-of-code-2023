package base

import (
	"bufio"
	"log"
	"os"
)

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

func ReadFileToCharacters(filePath string) [][]rune {
	lines := ReadFileToLines(filePath)
	characters := make([][]rune, len(lines))
	for y, line := range lines {
		if len(characters) < y {
			characters = append(characters, make([]rune, len(line)))
		}
		for x, char := range line {
			characters[y][x] = char
		}
	}
	return characters
}
