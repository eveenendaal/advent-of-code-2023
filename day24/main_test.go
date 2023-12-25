package day24

import "testing"

func TestPart1Test(t *testing.T) {
	result := Part1("test.txt", 7, 27)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Part1("input.txt", 200000000000000, 400000000000000)
	if result != 13910 {
		t.Errorf("Expected 13910, got %d", result)
	}
}

func TestPart2Test(t *testing.T) {
	result := Part2("test.txt")
	if result != 47 {
		t.Errorf("Expected 47, got %d", result)
	}
}
