package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

// Delta returns a new position from a row delta and col delta.
func (position Position) Delta(row, col int) Position {
	return Position{
		Row: position.Row + row,
		Col: position.Col + col,
	}
}

// NewPosition creates a new position.
func NewPosition(row, col int) Position {
	return Position{Row: row, Col: col}
}

// Move moves into a given direction and a certain number of times.
func (position Position) Move(direction Direction, moves int) Position {
	switch direction {
	case Up:
		return position.Delta(-moves, 0)
	case Down:
		return position.Delta(moves, 0)
	case Left:
		return position.Delta(0, -moves)
	case Right:
		return position.Delta(0, moves)
	}

	panic("not handled")
}

// Location represents a given position and direction.
type Location struct {
	Position  Position
	Direction Direction
}

// NewLocation creates a new location.
func NewLocation(row, col int, direction Direction) Location {
	return Location{
		Position:  NewPosition(row, col),
		Direction: direction,
	}
}

// Direction enum
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// Turn turns left or right.
func (d Direction) Turn(turn Direction) Direction {
	if turn != Left && turn != Right {
		panic("should be left or right")
	}

	switch d {
	case Up:
		return turn
	case Down:
		switch turn {
		case Left:
			return Right
		case Right:
			return Left
		}
	case Left:
		switch turn {
		case Left:
			return Down
		case Right:
			return Up
		}
	case Right:
		switch turn {
		case Left:
			return Up
		case Right:
			return Down
		}
	}

	panic("not handled")
}

// Position represents a given position (row/col)
type Position struct {
	Row int
	Col int
}

// Turn turns left or right.
func (location Location) Turn(direction Direction, moves int) Location {
	newDirection := location.Direction.Turn(direction)
	position := location.Position.Move(newDirection, moves)
	return Location{Position: position, Direction: newDirection}
}

// Straight moves in the current direction.
func (location Location) Straight(moves int) Location {
	pos := location.Position.Move(location.Direction, moves)
	return Location{Position: pos, Direction: location.Direction}
}

func shortest(board map[Position]int, target Position, minStraight, maxStraight int) int {
	type state struct {
		loc      Location
		straight int
	}
	type entry struct {
		state
		heatLoss int
	}

	queue := pq.NewWith(func(a, b any) int {
		p1 := a.(entry).heatLoss
		p2 := b.(entry).heatLoss
		return p1 - p2
	})

	queue.Enqueue(entry{
		state: state{
			loc:      NewLocation(0, 1, Right),
			straight: 1,
		},
	})
	queue.Enqueue(entry{
		state: state{
			loc:      NewLocation(1, 0, Down),
			straight: 1,
		},
	})
	visited := make(map[state]int)

	for !queue.Empty() {
		next, _ := queue.Dequeue()
		nextEntry := next.(entry)
		nextPosition := nextEntry.loc.Position

		if _, exists := board[nextPosition]; !exists {
			continue
		}

		newHeatLoss := board[nextPosition] + nextEntry.heatLoss
		if nextPosition == target {
			// Thanks to the priority queue, at this stage we already know this is the
			// shortest path.
			return newHeatLoss
		}

		if existingHeatLoss, exists := visited[nextEntry.state]; exists {
			if existingHeatLoss <= newHeatLoss {
				continue
			}
		}
		visited[nextEntry.state] = newHeatLoss

		if nextEntry.straight >= minStraight {
			queue.Enqueue(entry{
				state: state{
					loc:      nextEntry.loc.Turn(Left, 1),
					straight: 1,
				},
				heatLoss: newHeatLoss,
			})

			queue.Enqueue(entry{
				state: state{
					loc:      nextEntry.loc.Turn(Right, 1),
					straight: 1,
				},
				heatLoss: newHeatLoss,
			})
		}

		if nextEntry.straight < maxStraight {
			queue.Enqueue(entry{
				state: state{
					loc:      nextEntry.loc.Straight(1),
					straight: nextEntry.straight + 1,
				},
				heatLoss: newHeatLoss,
			})
		}
	}
	panic("no result found")
}

func Solve(filePath string, minStraight int, maxStraight int) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var board = make(map[Position]int)

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			value, _ := strconv.Atoi(string(char))
			board[Position{Row: y, Col: x}] = value
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// The target is the bottom right corner.
	maxX := 0
	maxY := 0
	for position, _ := range board {
		if position.Col > maxX {
			maxX = position.Col
		}
		if position.Row > maxY {
			maxY = position.Row
		}
	}

	target := Position{Row: maxY, Col: maxX}
	return shortest(board, target, minStraight, maxStraight)
}
