package day25

import (
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"strings"
)

type Edge struct {
	from, to string
}

func Part1(filePath string) int {
	lines := aoc.ReadFileToLines(filePath)
	nodes := getNodes(lines)

	// Remove 3 edges with the highest count
	var removedEdges []Edge
	for i := 0; i < 3; i++ {
		edgeCount := countEdges(nodes)
		edge := findAndRemoveMax(edgeCount)
		removeEdge(nodes, edge)
		removedEdges = append(removedEdges, edge)
	}

	// Count nodes in both groups
	countA := countNodes(nodes, removedEdges[0].from)
	countB := len(nodes) - countA

	// Multiply
	return countA * countB
}

func removeEdge(nodes map[string][]string, edge Edge) {
	newEdges := []string{}
	for _, val := range nodes[edge.from] {
		if val != edge.to {
			newEdges = append(newEdges, val)
		}
	}
	nodes[edge.from] = newEdges
	newEdges = []string{}
	for _, val := range nodes[edge.to] {
		if val != edge.from {
			newEdges = append(newEdges, val)
		}
	}
	nodes[edge.to] = newEdges
}

func findAndRemoveMax(edges map[Edge]int) Edge {
	currentMax := 0
	var maxEdge Edge
	for key, val := range edges {
		if val > currentMax {
			currentMax = val
			maxEdge = key
		}
	}
	delete(edges, maxEdge)
	return maxEdge
}

func countEdges(nodes map[string][]string) map[Edge]int {
	encountered := map[Edge]int{}
	for from := range nodes {
		walkNodes(nodes, from, encountered)
	}
	return encountered
}

func countNodes(nodes map[string][]string, start string) int {
	visited := map[string]bool{}
	queue := []string{start}

	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]

		for _, to := range nodes[from] {
			if _, found := visited[to]; found {
				continue
			}
			queue = append(queue, to)
			visited[to] = true
		}
	}
	return len(visited)
}

func walkNodes(nodes map[string][]string, start string, encountered map[Edge]int) {
	visited := map[string]bool{}
	queue := []string{start}

	for len(queue) > 0 {
		// pop first element
		from := queue[0]
		queue = queue[1:]

		// add all children to queue unless already visited
		for _, to := range nodes[from] {
			if _, found := visited[to]; found {
				continue
			}
			queue = append(queue, to)
			visited[to] = true
			// make sure we always have the same order
			var edge Edge
			if from < to {
				edge = Edge{from, to}
			} else {
				edge = Edge{to, from}
			}
			encountered[edge]++
		}
	}
}

func getNodes(lines []string) map[string][]string {
	nodes := map[string][]string{}
	for _, line := range lines {
		split := strings.Split(line, ": ")
		from := split[0]
		if _, inside := nodes[from]; !inside {
			nodes[from] = []string{}
		}

		to := strings.Fields(split[1])
		for _, target := range to {
			nodes[from] = append(nodes[from], target)
			if _, inside := nodes[target]; !inside {
				nodes[target] = []string{}
			}
			nodes[target] = append(nodes[target], from)
		}
	}
	return nodes
}
