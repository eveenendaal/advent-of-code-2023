package main

import "testing"

func TestPart1(t *testing.T) {
	result := Solve("test.txt", false)
	// assert the result is 6440
	if result != 6440 {
		t.Errorf("Expected 6440, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := Solve("test.txt", true)
	// assert the result is 5905
	if result != 5905 {
		t.Errorf("Expected 5905, got %d", result)
	}
}
