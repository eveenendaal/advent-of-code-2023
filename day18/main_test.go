package main

import "testing"

func TestPart1Test(t *testing.T) {
	result := Solve("test.txt")
	if result != 62 {
		t.Errorf("Expected 62, got %d", result)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Solve("input.txt")
	if result != 62 {
		t.Errorf("Expected 62, got %d", result)
	}
}

func TestFunction3(t *testing.T) {
	// TODO: Implement test
}

func TestFunction4(t *testing.T) {
	// TODO: Implement test
}
