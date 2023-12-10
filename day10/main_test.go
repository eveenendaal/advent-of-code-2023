package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("test1.txt")
	if result != 4 {
		t.Errorf("Expected result to be 4, got %d", result)
	}

	result = Part1("test2.txt")
	if result != 8 {
		t.Errorf("Expected result to be 8, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Write test cases for getPreviousValue function
}
