package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func hash(input string) int {
	currentValue := 0

	// iterate over each character
	for _, char := range input {
		// convert to int
		charValue := int(char)
		// add to current value
		currentValue += charValue
		// multiply by 17
		currentValue *= 17
		// mod by 256
		currentValue %= 256
	}

	return currentValue
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	total := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// split string on ,
		blocks := strings.Split(scanner.Text(), ",")
		for _, block := range blocks {
			total += hash(block)
		}
	}

	err = file.Close()
	if err != nil {
		log.Fatalf("failed to close file: %s", err)
	}

	return total
}

func main() {
	fmt.Printf("Part 1: %d\n", Part1("input.txt"))
}
