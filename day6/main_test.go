package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("test.txt")
	// assert the result is 288
	if result != 288 {
		t.Errorf("Expected 288, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Write test cases for main function
}
