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
	answer := int64(11687500)
	diff := result - answer
	if result != answer {
		t.Errorf("Expected %d, got %d (diff %d)", answer, result, diff)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Part1("input.txt", 1000)
	answer := int64(929810733)
	diff := result - answer
	if result != answer {
		t.Errorf("Expected %d, got %d (diff %d)", answer, result, diff)
	}
}

func TestPart2Solution(t *testing.T) {
	result := Part2("input.txt", "rx")
	if result != 231657829136023 {
		t.Errorf("Expected 231657829136023, got %d", result)
	}
}
