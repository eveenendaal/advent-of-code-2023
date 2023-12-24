package day23

import (
	"fmt"
	"testing"
)

func TestPart1Test(t *testing.T) {
	expected := 94
	actual := Part1("test.txt")
	if actual != expected {
		t.Errorf("Solve() = %d, expected %d", actual, expected)
	} else {
		fmt.Printf("Part1: %d\n", actual)
	}
}

func TestPart1Solution(t *testing.T) {
	expected := 2254
	actual := Part1("input.txt")
	if actual != expected {
		t.Errorf("Solve() = %d, expected %d", actual, expected)
	} else {
		fmt.Printf("Part1: %d\n", actual)
	}
}

func TestPart2Test(t *testing.T) {
	expected := 154
	actual := Part2("test.txt")
	if actual != expected {
		t.Errorf("Solve() = %d, expected %d", actual, expected)
	} else {
		fmt.Printf("Part1: %d\n", actual)
	}
}

func TestPart2Solution(t *testing.T) {
	expected := 6394
	actual := Part2("input.txt")
	if actual != expected {
		t.Errorf("Solve() = %d, expected %d", actual, expected)
	} else {
		fmt.Printf("Part1: %d\n", actual)
	}
}
