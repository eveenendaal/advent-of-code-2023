package day24

import "testing"

func TestPart1Test(t *testing.T) {
	result := Part1("test.txt")
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Part1("input.txt")
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
