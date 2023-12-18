package main

import "testing"

func TestPart1Test(t *testing.T) {
	result := Part1("test.txt")
	if result != 62 {
		t.Errorf("Expected 62, got %d", result)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Part1("input.txt")
	if result != 62 {
		t.Errorf("Expected 62, got %d", result)
	}
}

func TestPart2Test(t *testing.T) {
	result := Part2("test.txt")
	answer := 952408144115
	if result != answer {
		diff := answer - result
		t.Errorf("Expected %d, got %d (diff %d)", answer, result, diff)
	}
}

func TestPart2Solution(t *testing.T) {
	result := Part2("input.txt")
	answer := 93325849869340
	if result != answer {
		diff := answer - result
		t.Errorf("Expected %d, got %d (diff %d)", answer, result, diff)
	}
}
