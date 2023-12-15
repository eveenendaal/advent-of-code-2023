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
	for _, c := range column {
		fmt.Printf("%s", string(c))
	}
	fmt.Println()
}

func Part1(filePath string) int {
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

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

func sortColumns(data [][]rune, direction int) {
	rows := len(data)
	columns := len(data[0])

	switch direction {
	case NORTH:
		for i := 0; i < columns; i++ {
			for j := 0; j < rows; j++ {
				if data[j][i] == 'O' {
					for k := j; k > 0; k-- {
						if data[k-1][i] == '#' {
							break
						}
						data[k-1][i], data[k][i] = data[k][i], data[k-1][i]
					}
				}
			}
		}
	case EAST:
		for i := 0; i < rows; i++ {
			for j := columns - 1; j >= 0; j-- {
				if data[i][j] == 'O' {
					for k := j; k < columns-1; k++ {
						if data[i][k+1] == '#' {
							break
						}
						data[i][k+1], data[i][k] = data[i][k], data[i][k+1]
					}
				}
			}
		}
	case WEST:
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				if data[i][j] == 'O' {
					for k := j; k > 0; k-- {
						if data[i][k-1] == '#' {
							break
						}
						data[i][k-1], data[i][k] = data[i][k], data[i][k-1]
					}
				}
			}
		}
	case SOUTH:
		for i := 0; i < columns; i++ {
			for j := rows - 1; j >= 0; j-- {
				if data[j][i] == 'O' {
					for k := j; k < rows-1; k++ {
						if data[k+1][i] == '#' {
							continue
						}
						data[k+1][i], data[k][i] = data[k][i], data[k+1][i]
					}
				}
			}
		}
	}
}

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	rows := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// split string into runes
		rows = append(rows, []rune(line))
	}
	fmt.Println()

	cycles := 1000000000

	for i := 0; i < cycles; i++ {
		sortColumns(rows, NORTH)
		sortColumns(rows, WEST)
		sortColumns(rows, SOUTH)
		sortColumns(rows, EAST)

		if i%1000000 == 0 {
			fmt.Printf("Iteration: %d of %v\n", i, float32(i)/float32(cycles)*100)
		}
	}

	// Sort the columns
	total := 0
	for i, row := range rows {
		for _, c := range row {
			if c == 'O' {
				total += len(rows) - i
			}
		}
	}

	err = file.Close()
	if err != nil {
		log.Fatalf("failed to close file: %s", err)
	}

	return total
}

func main() {
	// fmt.Println("Part 1 Solution: ", Part1("input.txt"))
	fmt.Println("Part 2 Solution: ", Part2("input.txt"))
}
