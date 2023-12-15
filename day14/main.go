package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func printColumns(columns [][]rune) {
	for _, column := range columns {
		printColumn(column)
	}
	fmt.Println()
}

func printColumn(column []rune) {
	fmt.Printf("Column: ")
	for _, c := range column {
		fmt.Printf("%s", string(c))
	}
	fmt.Println()
}

func Solve(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	total := 0
	columns := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		for columnNumber, c := range line {
			if len(columns) <= columnNumber {
				columns = append(columns, make([]rune, 0))
			}
			columns[columnNumber] = append(columns[columnNumber], c)
		}
	}
	fmt.Println()

	printColumns(columns)

	// Sort the columns
	for _, column := range columns {
		// move all the O's to the left until they run into a #
		for i := 0; i < len(column); i++ {
			if column[i] == 'O' {
				for j := i; j > 0; j-- {
					if column[j-1] == '#' {
						break
					}
					column[j-1], column[j] = column[j], column[j-1]
				}
			}
		}

		printColumn(column)
		for i := 0; i < len(column); i++ {
			if column[i] == 'O' {
				total += len(column) - i
			}
		}
	}

	// printColumns(columns)

	err = file.Close()
	if err != nil {
		log.Fatalf("failed to close file: %s", err)
	}

	return total
}

func main() {
	fmt.Println("Part 1 Solution: ", Solve("input.txt"))
}
