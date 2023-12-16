package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	Up = iota
	Right
	Down
	Left
)

type Beam struct {
	x         int
	y         int
	direction int
	number    int
	duplicate bool
}

type Point struct {
	x                int
	y                int
	char             rune
	activeDirections []int
}

func (b *Beam) move() {
	switch b.direction {
	case Up:
		b.y--
	case Right:
		b.x++
	case Down:
		b.y++
	case Left:
		b.x--
	}
}

func directionToString(direction int) string {
	switch direction {
	case Up:
		return "Up"
	case Right:
		return "Right"
	case Down:
		return "Down"
	case Left:
		return "Left"
	default:
		return "Unknown"
	}
}

func handleBeam(startBeam *Beam, grid [][]Point) {
	var beams = make([]Beam, 0)
	beams = append(beams, *startBeam)
	beamCounter := 1

	for {
		newBeams := make([]Beam, 0)

		// Update the direction
		for i, beam := range beams {
			point := grid[beam.y][beam.x]

			// Check if we've already been here
			for _, direction := range point.activeDirections {
				if direction == beam.direction {
					// This beam is a duplicate
					beams[i].duplicate = true
					continue
				}
			}
			grid[beam.y][beam.x].activeDirections = append(point.activeDirections, beam.direction)

			switch point.char {
			case '|':
				if beam.direction == Left || beam.direction == Right {
					newBeams = append(newBeams, Beam{int(beam.x), int(beam.y), Up, beamCounter, false})
					beamCounter++
					beams[i].direction = Down
				}
			case '-':
				if beam.direction == Up || beam.direction == Down {
					newBeams = append(newBeams, Beam{int(beam.x), int(beam.y), Left, beamCounter, false})
					beamCounter++
					beams[i].direction = Right
				}
			case '/':
				switch beam.direction {
				case Up:
					beams[i].direction = Right
				case Right:
					beams[i].direction = Up
				case Down:
					beams[i].direction = Left
				case Left:
					beams[i].direction = Down
				}
			case '\\':
				switch beam.direction {
				case Up:
					beams[i].direction = Left
				case Right:
					beams[i].direction = Down
				case Down:
					beams[i].direction = Right
				case Left:
					beams[i].direction = Up
				}
			}
			// Move the beam forward
			// fmt.Printf("Beam %d: %d, %d, %s\n", beam.number, beam.x, beam.y, directionToString(beam.direction))
			beams[i].move()
		}

		// Remove beams that have left the grid
		for i := len(beams) - 1; i >= 0; i-- {
			if beams[i].x < 0 || beams[i].y < 0 || beams[i].x >= len(grid[0]) || beams[i].y >= len(grid) {
				// fmt.Printf("Beam left grid %d\n", beams[i].number)
				beams = append(beams[:i], beams[i+1:]...)
			} else if beams[i].duplicate {
				// fmt.Printf("Beam repeat %d\n", beams[i].number)
				beams = append(beams[:i], beams[i+1:]...)
			}
		}

		// Add new beams to the list
		if len(newBeams) > 0 {
			beams = append(beams, newBeams...)
		}

		if len(beams) == 0 {
			break
		}
	}

}

func getTotal(grid [][]Point) int {
	total := 0
	for _, row := range grid {
		for _, point := range row {
			if len(point.activeDirections) > 0 {
				total++
			}
		}
	}
	return total
}

func Solve(filePath string, part1 bool) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid = make([][]Point, 0)

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		if len(grid) <= y {
			grid = append(grid, make([]Point, 0))
		}

		for x, c := range scanner.Text() {
			grid[y] = append(grid[y], Point{x, y, c, make([]int, 0)})
		}
		y++
	}

	total := 0

	if part1 {
		beam := Beam{0, 0, Right, 0, false}
		handleBeam(&beam, grid)
		total = getTotal(grid)
	} else {
		// Part 2

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func main() {
	// fmt.Printf("Part 1: %d\n", Solve("input.txt", true))
	fmt.Printf("Part 2: %d\n", Solve("input.txt", false))
}
