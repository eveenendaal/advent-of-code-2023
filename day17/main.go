package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Solve(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	var grid = make([][]int, 0)

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		if len(grid) <= y {
			grid = append(grid, []int{})
		}
		line := scanner.Text()
		for _, char := range line {
			value, _ := strconv.Atoi(string(char))
			grid[y] = append(grid[y], value)
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}
