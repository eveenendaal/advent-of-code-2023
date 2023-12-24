package day24

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"strconv"
	"strings"
)

type Hailstone struct {
	px int
	py int
	pz int

	vx int
	vy int
	vz int
}

func NewHailstone(line string) *Hailstone {
	// split on @
	parts := strings.Split(line, "@")
	// Parse position
	positionParts := strings.Split(parts[0], ",")
	px, _ := strconv.Atoi(strings.TrimSpace(positionParts[0]))
	py, _ := strconv.Atoi(strings.TrimSpace(positionParts[1]))
	pz, _ := strconv.Atoi(strings.TrimSpace(positionParts[2]))

	// Parse velocity
	velocityParts := strings.Split(parts[1], ",")
	vx, _ := strconv.Atoi(strings.TrimSpace(velocityParts[0]))
	vy, _ := strconv.Atoi(strings.TrimSpace(velocityParts[1]))
	vz, _ := strconv.Atoi(strings.TrimSpace(velocityParts[2]))

	// Create Hailstone
	return &Hailstone{
		px: px,
		py: py,
		pz: pz,
		vx: vx,
		vy: vy,
		vz: vz,
	}
}

func Part1(filePath string) int {
	lines := aoc.ReadFileToLines(filePath)

	for _, line := range lines {
		hailstone := NewHailstone(line)
		fmt.Printf("Hailstone: %v\n", hailstone)
	}

	return 0
}
