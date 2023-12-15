package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	total := 0
	columns := make([][]rune, 0)
	rowNumber := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
		// Move each character into a column
		for i, c := range line {
			if len(columns) <= i {
				columns = append(columns, make([]rune, len(line)))
			}
			columns[i][rowNumber] = c
		}

		rowNumber++
	}

	for _, column := range columns {
		fmt.Printf("Column: ")
		for _, c := range column {
			fmt.Printf("%s", string(c))
		}
		fmt.Println()

		// move all the O's to the left until they run into a #
		for i := 0; i < len(column); i++ {
			if column[i] == 'O' {
				for j := i; j > 0; j-- {
					if column[j-1] == '#' {
						continue
					}
					columns[i][j] = column[j-1]
					columns[i][j-1] = 'O'
				}
			}
		}
	}

	err = file.Close()
	if err != nil {
		log.Fatalf("failed to close file: %s", err)
	}

	return total
}
