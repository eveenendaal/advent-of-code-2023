package main

import (
	"testing"
)

func TestPart1Test(t *testing.T) {
	result := Part1("test.txt")
	if result != 19114 {
		t.Errorf("Expected 19114, got %d", result)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Part1("input.txt")
	if result != 434147 {
		t.Errorf("Expected 434147, got %d", result)
	}
}

func TestPart2Test(t *testing.T) {
	result := Part2("test.txt")
	if result != 167409079868000 {
		t.Errorf("Expected 167409079868000, got %d", result)
	}
}

func TestPart2Solution(t *testing.T) {
	result := Part2("input.txt")
	if result != 136146366355609 {
		t.Errorf("Expected 136146366355609, got %d", result)
	}
}
