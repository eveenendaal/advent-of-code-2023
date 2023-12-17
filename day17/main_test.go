package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	result := Solve("test.txt", 0, 3)
	if result != 102 {
		t.Errorf("Expected 102, got %d", result)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Solve("input.txt", 0, 3)
	if result != 814 {
		t.Errorf("Expected 814, got %d", result)
	}
}

func TestPart2Test1(t *testing.T) {
	result := Solve("test.txt", 4, 10)
	if result != 94 {
		t.Errorf("Expected 94, got %d", result)
	}
}

func TestPart2Test2(t *testing.T) {
	result := Solve("test2.txt", 4, 10)
	if result != 71 {
		t.Errorf("Expected 71, got %d", result)
	}
}

func TestPart2Solution(t *testing.T) {
	result := Solve("input.txt", 4, 10)
	fmt.Printf("Result: %d\n", result)
	if result != 974 {
		t.Errorf("Expected 974, got %d", result)
	}
}
