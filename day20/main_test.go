package main

import "testing"

func TestPart1Test1a(t *testing.T) {
	result := Part1("test1.txt", 1)
	if result != 32 {
		t.Errorf("Expected 32, got %d", result)
	}

}

func TestPart1Test1b(t *testing.T) {
	result := Part1("test1.txt", 1000)
	if result != 32000000 {
		t.Errorf("Expected 32000000, got %d", result)
	}
}

func TestPart1Test2(t *testing.T) {
	result := Part1("test2.txt", 1000)
	if result != 11687500 {
		t.Errorf("Expected 11687500, got %d", result)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Part1("input.txt", 1000)
	if result != 6658637668 {
		t.Errorf("Expected 6658637668, got %d", result)
	}
}
