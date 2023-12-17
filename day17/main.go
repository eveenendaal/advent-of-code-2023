package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Point struct {
	X int
	Y int
}

type VisitedWithSteps struct {
	node      map[int]int
	steps     int
	direction Direction
}

type Direction int

const (
	Up    = Direction(0)
	Down  = Direction(1)
	Left  = Direction(2)
	Right = Direction(3)
)

//type Item struct {
//	value    Point // The value of the item; arbitrary.
//	priority int   // The priority of the item in the queue.
//	// The index is needed by update and is maintained by the heap.Interface methods.
//	index      int // The index of the item in the heap.
//	history    []Point
//	directions []int
//}
//
//type Neighbor struct {
//	point     Point
//	direction int
//}
//
//type PriorityQueue []*Item
//
//func (pq PriorityQueue) Len() int { return len(pq) }
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
//	return pq[i].priority < pq[j].priority
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].index = i
//	pq[j].index = j
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	n := len(*pq)
//	item := x.(*Item)
//	item.index = n
//	*pq = append(*pq, item)
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	item := old[n-1]
//	old[n-1] = nil  // avoid memory leak
//	item.index = -1 // for safety
//	*pq = old[0 : n-1]
//	return item
//}
//
//func processGrid(grid [][]int, start Point, end Point) int {
//	// Initialize distance map and priority queue
//	dist := make(map[Point]int)
//	for y, row := range grid {
//		for x, _ := range row {
//			point := Point{x, y}
//			if point == start {
//				dist[point] = grid[start.Y][start.X]
//			} else {
//				dist[point] = math.MaxInt32
//			}
//		}
//	}
//
//	pq := make(PriorityQueue, len(dist))
//	i := 0
//	for point, distance := range dist {
//		pq[i] = &Item{
//			value:      point,
//			priority:   distance,
//			index:      i,
//			history:    []Point{},
//			directions: []int{},
//		}
//		i++
//	}
//	heap.Init(&pq)
//
//	// Dijkstra's algorithm
//	for pq.Len() > 0 {
//		item := heap.Pop(&pq).(*Item)
//		point := item.value
//
//		if point == end {
//			break
//		}
//
//		// Visit all neighbors
//		for _, neighbor := range getNeighbors(point, grid, item.history, item.directions) {
//			alt := dist[point] + neighbor.point.getCost(grid)
//			if alt < dist[neighbor.point] {
//				dist[neighbor.point] = alt
//
//				heap.Push(&pq, &Item{
//					value:      neighbor.point,
//					priority:   alt,
//					history:    append(item.history, point),
//					directions: append(item.directions, neighbor.direction),
//				})
//			}
//		}
//	}
//
//	return dist[end]
//}
//
//func getNeighbors(point Point, grid [][]int, history []Point, directions []int) []Neighbor {
//	row := grid[point.Y]
//	neighbors := make([]Neighbor, 0)
//	if point.X > 0 {
//		neighbors = append(neighbors, Neighbor{Point{point.X - 1, point.Y}, Left})
//	}
//	if point.X < len(row)-1 {
//		neighbors = append(neighbors, Neighbor{Point{point.X + 1, point.Y}, Right})
//	}
//	if point.Y > 0 {
//		neighbors = append(neighbors, Neighbor{Point{point.X, point.Y - 1}, Up})
//	}
//	if point.Y < len(grid)-1 {
//		neighbors = append(neighbors, Neighbor{Point{point.X, point.Y + 1}, Down})
//	}
//	// remove points in history
//	for _, h := range history {
//		for i, n := range neighbors {
//			if n.point == h {
//				neighbors = append(neighbors[:i], neighbors[i+1:]...)
//				break
//			}
//		}
//	}
//	// remove points are the fourth in a row
//	for _, n := range neighbors {
//		// the third from last point in history
//		if len(directions) > 2 {
//			recentDirections := directions[len(directions)-3:]
//			// if n.direction and all recentDirections are the same, remove n
//			if n.direction == recentDirections[0] && n.direction == recentDirections[1] && n.direction == recentDirections[2] {
//				for i, neighbor := range neighbors {
//					if neighbor == n {
//						neighbors = append(neighbors[:i], neighbors[i+1:]...)
//						break
//					}
//				}
//			}
//		}
//	}
//
//	return neighbors
//}
//
//func (p Point) getCost(grid [][]int) int {
//	return grid[p.Y][p.X]
//}

func Solve(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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

	start := Point{0, 0}
	end := Point{len(grid[0]) - 1, len(grid) - 1}
	fmt.Printf("Start: %v\n", start)
	fmt.Printf("End: %v\n", end)
	//return processGrid(grid, start, end)

	// Answers
	// Part1: 814
	// Part2: 974
}

func main() {
	fmt.Printf("Part 1: %d\n", Solve("input.txt"))
}
