package main

import "testing"

func TestPart1Example1(t *testing.T) {
	result := Solve("test1.txt", false)
	if result != 52 {
		t.Errorf("Expected 52, got %d", result)
	}
}

func TestPart1Example2(t *testing.T) {
	result := Solve("test2.txt", false)
	if result != 1320 {
		t.Errorf("Expected 1320, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := Solve("test2.txt", true)
	if result != 145 {
		t.Errorf("Expected 145, got %d", result)
	}
}
