package main

import "testing"

func TestPart1Example1(t *testing.T) {
	result := Part1("test1.txt")
	if result != 52 {
		t.Errorf("Expected 52, got %d", result)
	}
}

func TestPart1Example2(t *testing.T) {
	result := Part1("test2.txt")
	if result != 1320 {
		t.Errorf("Expected 1320, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Add your test logic here
}
