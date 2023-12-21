package day21

import "testing"

func TestPart1Test(t *testing.T) {
	result := Part1("test.txt", 6)
	if result != 16 {
		t.Errorf("Expected result to be 16, got %d", result)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Part1("input.txt", 64)
	if result != 3746 {
		t.Errorf("Expected result to be 3746, got %d", result)
	}
}
