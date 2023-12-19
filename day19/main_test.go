package main

import "testing"

func TestPart1Test(t *testing.T) {
	result := Part1("test.txt")
	if result != 19114 {
		t.Errorf("Expected 19114, got %d", result)
	}
}

func TestPart2Solution(t *testing.T) {
	// TODO: Implement test
}
