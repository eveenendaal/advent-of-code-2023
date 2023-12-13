package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("test.txt")
	if result != 405 {
		t.Errorf("Expected %d, got %d", 405, result)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Write test case for another example scenario
}
