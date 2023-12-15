package main

import "testing"

// func TestPart1(t *testing.T) {
// 	result := Part1("test.txt")
// 	if result != 136 {
// 		t.Errorf("Expected 136, got %d", result)
// 	}
// }

func TestPart2(t *testing.T) {
	result := Part2("test.txt")
	if result != 64 {
		t.Errorf("Expected 64, got %d", result)
	}
}
