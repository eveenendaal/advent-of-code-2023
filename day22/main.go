package day22

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"strconv"
	"strings"
)

type Brick struct {
	x   int
	y   int
	z   int
	row int
}

func Part1(filePath string) int {
	lines := aoc.ReadFileToLines(filePath)
	bricks := make([]Brick, 0)

	for row, line := range lines {
		// splits on ~
		parts := strings.Split(line, "~")
		for _, part := range parts {
			// splits on ,
			coordinates := strings.Split(part, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])
			z, _ := strconv.Atoi(coordinates[2])
			brick := Brick{x, y, z, row}
			bricks = append(bricks, brick)
		}
	}

	for _, brick := range bricks {
		fmt.Printf("brick: %v\n", brick)
	}

	return 0
}
