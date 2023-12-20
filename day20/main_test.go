package main

import "testing"

func TestPart1Test1(t *testing.T) {
	result := Part1("test1.txt")
	if result != 32000000 {
		t.Errorf("Expected 32000000, got %d", result)
	}
}

func TestPart1Test2(t *testing.T) {
	result := Part1("test2.txt")
	if result != 11687500 {
		t.Errorf("Expected 11687500, got %d", result)
	}
}
