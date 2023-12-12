package main

import "testing"

func TestPart1(t *testing.T) {
	result := Solve("test.txt", false)
	if result != 21 {
		t.Errorf("Expected 21, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// result := Solve("test.txt", true)
	// if result != 525152 {
	// 	t.Errorf("Expected 525152, got %d", result)
	// }
}
