package day22

import (
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"slices"
	"strconv"
	"strings"
)

func Part1(filename string) int {
	input := aoc.ReadFileToLines(filename)
	sortedBricks := getBricks(input)
	getNumberOfBricksShiftedAfterSettling(sortedBricks)
	ans1, _ := findBricksToDisintegrate(sortedBricks)
	return ans1
}

func Part2(filename string) int {
	input := aoc.ReadFileToLines(filename)
	sortedBricks := getBricks(input)
	getNumberOfBricksShiftedAfterSettling(sortedBricks)
	_, ans2 := findBricksToDisintegrate(sortedBricks)
	return ans2
}

type Brick struct {
	start Point
	end   Point
}

type Point struct {
	x int
	y int
	z int
}

func getBricks(input []string) []Brick {
	var bricks []Brick
	for _, line := range input {
		parts := strings.Split(line, "~")
		var digits []int
		for _, part := range parts {
			digitStrs := strings.Split(part, ",")
			for _, digitStr := range digitStrs {
				digit, _ := strconv.Atoi(digitStr)
				digits = append(digits, digit)
			}

		}
		bricks = append(bricks, Brick{
			start: Point{digits[0], digits[1], digits[2]},
			end:   Point{digits[3], digits[4], digits[5]},
		})
	}

	// sort bricks by z coordinate
	slices.SortFunc(bricks, func(a, b Brick) int {
		return a.start.z - b.start.z
	})
	return bricks
}

func findNewZCoordinates(xyPlane [][]int, p1, p2 Point) (int, int) {
	z1, z2 := 0, 0

	// if x and z coordinates are same, Brick elongates in y dir
	// note: this condition will also handle if all coordinates are same, ie 1 cubic meter
	if p1.x == p2.x && p1.z == p2.z {
		for i := p1.y; i <= p2.y; i++ {
			if xyPlane[p1.x][i] > z1 {
				z1 = xyPlane[p1.x][i]
			}
		}

		// max value only needs to be incremented by one
		// as Brick has only width 1 in z dir
		z1, z2 = z1+1, z1+1

		for i := p1.y; i <= p2.y; i++ {
			xyPlane[p1.x][i] = z1
		}

		// if y and z coordinates are same, Brick elongates in x dir
	} else if p1.y == p2.y && p1.z == p2.z {
		for i := p1.x; i <= p2.x; i++ {
			if xyPlane[i][p1.y] > z1 {
				z1 = xyPlane[i][p1.y]
			}
		}

		// max value only needs to be incremented by one
		// as Brick has only width 1 in z dir
		z1, z2 = z1+1, z1+1

		// value only needs to be updated by one
		// as Brick has only width 1 in z dir
		for i := p1.x; i <= p2.x; i++ {
			xyPlane[i][p1.y] = z1
		}

		// if x and y coordinates are same, Brick elongates in z dir
	} else if p1.x == p2.x && p1.y == p2.y {
		// as the Brick extends in Z dir, only a single Brick in
		// the count plane will need to be updated with latest Z coordinate
		z1 = xyPlane[p1.x][p1.y] + 1
		z2 = z1 + p2.z - p1.z
		xyPlane[p1.x][p1.y] = z2
	}

	return z1, z2
}

func findBricksToDisintegrate(bricks []Brick) (int, int) {
	allowed := 0
	sum := 0

	for i, _ := range bricks {
		newBrick := make([]Brick, len(bricks))
		copy(newBrick, bricks)
		if i == 0 {
			newBrick = newBrick[1:]
		} else if i == len(newBrick)-1 {
			newBrick = bricks[:len(newBrick)-1]
		} else {
			newBrick = append(newBrick[:i], newBrick[i+1:]...)
		}

		changes := getNumberOfBricksShiftedAfterSettling(newBrick)
		if changes == 0 {
			allowed++
		}
		sum += changes

	}
	return allowed, sum
}

func getNumberOfBricksShiftedAfterSettling(bricks []Brick) int {
	count := 0

	xyPlane := make([][]int, 10)
	for i, _ := range xyPlane {
		xyPlane[i] = make([]int, 10)
	}

	for i, b := range bricks {

		b.start.z, b.end.z = findNewZCoordinates(xyPlane, b.start, b.end)

		if b.start != bricks[i].start || b.end != bricks[i].end {
			bricks[i] = Brick{b.start, b.end}
			count++
		}
	}
	return count
}
