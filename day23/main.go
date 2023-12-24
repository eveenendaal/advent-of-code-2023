package day23

import (
	_ "embed"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"math"
	"slices"
)

type PosCost struct {
	pos  aoc.Position
	cost int
}

type Grid map[aoc.Position]uint8

func GridBounds(grid map[aoc.Position]uint8) (minX, maxX, minY, maxY int) {
	minX, maxX = math.MaxInt, math.MinInt
	minY, maxY = math.MaxInt, math.MinInt
	for p := range grid {
		minX = min(p.Col, minX)
		maxX = max(p.Col, maxX)
		minY = min(p.Row, minY)
		maxY = max(p.Row, maxY)
	}
	return minX, maxX, minY, maxY
}

func exploreSinglePath(grid Grid, previous aoc.Position, current aoc.Position, cost int, part2 bool) (PosCost, bool) {
	if c, ok := grid[current]; ok && c != '#' {
		var cpt int
		for _, ne := range Neighbors4(current) {
			if c, ok := grid[ne]; ok && c != '#' {
				cpt++
			}
		}
		if cpt > 2 {
			return PosCost{pos: current, cost: cost}, true
		}
	}

	if !part2 {
		// cut branches in part1
		if c, ok := grid[current]; ok && c != '.' {
			if current.Col > previous.Col && c != '>' ||
				current.Col < previous.Col && c != '<' ||
				current.Row > previous.Row && c != 'v' ||
				current.Row < previous.Row && c != '^' {
				return PosCost{}, false
			}
		}
	}

	for _, n := range Neighbors4(current) {
		if c, ok := grid[n]; ok && c != '#' && n != previous {
			return exploreSinglePath(grid, current, n, cost+1, part2)
		}
	}

	return PosCost{pos: current, cost: cost}, true
}

func explore(neighbors Graph, point, goal aoc.Position, visited map[aoc.Position]bool, cost int, maxCost int) int {
	if point == goal {
		if cost > maxCost {
			maxCost = cost
		}
		return maxCost
	}

	visited[point] = true
	for _, pc := range neighbors[point] {
		if !visited[pc.pos] {
			maxCost = explore(neighbors, pc.pos, goal, visited, cost+pc.cost, maxCost)
		}
	}
	visited[point] = false
	return maxCost
}

type Graph map[aoc.Position][]PosCost

func Neighbors4(p aoc.Position) []aoc.Position {
	return []aoc.Position{
		{Col: p.Col, Row: p.Row - 1},
		{Col: p.Col, Row: p.Row + 1},
		{Col: p.Col - 1, Row: p.Row},
		{Col: p.Col + 1, Row: p.Row},
	}
}

func buildGraph(grid Grid, start aoc.Position, part2 bool) Graph {
	var res = make(map[aoc.Position][]PosCost)

	var todo = []aoc.Position{}
	todo = append(todo, start)

	for len(todo) > 0 {
		p := todo[0]
		todo = todo[1:]
		if c, ok := grid[p]; !ok || c == '#' {
			continue
		}
		for _, n := range Neighbors4(p) {
			if c, ok := grid[n]; !ok || c == '#' {
				continue
			}
			pc, ok := exploreSinglePath(grid, p, n, 1, part2)
			if ok && !slices.Contains(res[p], pc) {
				res[p] = append(res[p], pc)
				todo = append(todo, pc.pos)
			}
		}
	}

	return res
}

func solve(filepath string, part2 bool) int {
	grid := make(map[aoc.Position]uint8)
	characters := aoc.ReadFileToCharacters(filepath)
	for y, row := range characters {
		for x, c := range row {
			grid[aoc.Position{Col: x, Row: y}] = uint8(c)
		}
	}

	minX, maxX, minY, maxY := GridBounds(grid)
	start := aoc.Position{Col: minX + 1, Row: minY}
	end := aoc.Position{Col: maxX - 1, Row: maxY}

	neighbors := buildGraph(grid, start, part2)

	var goal = end
	var path = 0

	if part2 && len(neighbors[end]) > 0 {
		// skip last path
		goal = neighbors[end][0].pos
		path = neighbors[end][0].cost
	}

	visited := make(map[aoc.Position]bool)
	return explore(neighbors, start, goal, visited, path, 0)
}

func Part1(filPath string) int {
	return solve(filPath, false)
}

func Part2(filPath string) int {
	return solve(filPath, true)
}
