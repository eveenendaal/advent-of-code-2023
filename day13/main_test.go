package main

import "testing"

func TestPart1(t *testing.T) {
	result := Solve("test.txt", false)
	if result != 405 {
		t.Errorf("Expected %d, got %d", 405, result)
	}
}

func TestPart2(t *testing.T) {
	result := Solve("test.txt", true)
	if result != 400 {
		t.Errorf("Expected %d, got %d", 400, result)
	}
}
