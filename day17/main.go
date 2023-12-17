package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

// Delta returns a new position from a row delta and col delta.
func (p Position) Delta(row, col int) Position {
	return Position{
		Row: p.Row + row,
		Col: p.Col + col,
	}
}

// NewPosition creates a new position.
func NewPosition(row, col int) Position {
	return Position{Row: row, Col: col}
}

// Move moves into a given direction and a certain number of times.
func (p Position) Move(direction Direction, moves int) Position {
	switch direction {
	case Up:
		return p.Delta(-moves, 0)
	case Down:
		return p.Delta(moves, 0)
	case Left:
		return p.Delta(0, -moves)
	case Right:
		return p.Delta(0, moves)
	}

	panic("not handled")
}

// Location represents a given position and direction.
type Location struct {
	Pos Position
	Dir Direction
}

// NewLocation creates a new location.
func NewLocation(row, col int, dir Direction) Location {
	return Location{
		Pos: NewPosition(row, col),
		Dir: dir,
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

// Rev reverses the current direction.
func (d Direction) Rev() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	panic("not handled")
}

// Rev moves in the reverse direction.
func (l Location) Rev(moves int) Location {
	dir := l.Dir.Rev()
	pos := l.Pos.Move(dir, moves)
	return Location{Pos: pos, Dir: dir}
}

// Turn turns left or right.
func (l Location) Turn(d Direction, moves int) Location {
	dir := l.Dir.Turn(d)
	pos := l.Pos.Move(dir, moves)
	return Location{Pos: pos, Dir: dir}
}

// Straight moves in the current direction.
func (l Location) Straight(moves int) Location {
	pos := l.Pos.Move(l.Dir, moves)
	return Location{Pos: pos, Dir: l.Dir}
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

	q := pq.NewWith(func(a, b any) int {
		p1 := a.(entry).heatLoss
		p2 := b.(entry).heatLoss
		return p1 - p2
	})

	q.Enqueue(entry{
		state: state{
			loc:      NewLocation(0, 1, Right),
			straight: 1,
		},
	})
	q.Enqueue(entry{
		state: state{
			loc:      NewLocation(1, 0, Down),
			straight: 1,
		},
	})
	visited := make(map[state]int)

	for !q.Empty() {
		t, _ := q.Dequeue()
		e := t.(entry)
		pos := e.loc.Pos

		if _, exists := board[pos]; !exists {
			continue
		}

		heat := board[pos] + e.heatLoss
		if pos == target {
			// Thanks to the priority queue, at this stage we already know this is the
			// shortest path.
			return heat
		}

		if v, exists := visited[e.state]; exists {
			if v <= heat {
				continue
			}
		}
		visited[e.state] = heat

		if e.straight >= minStraight {
			q.Enqueue(entry{
				state: state{
					loc:      e.loc.Turn(Left, 1),
					straight: 1,
				},
				heatLoss: heat,
			})

			q.Enqueue(entry{
				state: state{
					loc:      e.loc.Turn(Right, 1),
					straight: 1,
				},
				heatLoss: heat,
			})
		}

		if e.straight < maxStraight {
			q.Enqueue(entry{
				state: state{
					loc:      e.loc.Straight(1),
					straight: e.straight + 1,
				},
				heatLoss: heat,
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
