package main

import "testing"

func TestPart1(t *testing.T) {
	result := Solve("test.txt")
	if result != 102 {
		t.Errorf("Expected 102, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Implement test
}
