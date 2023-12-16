package main

import "testing"

func TestPart1(t *testing.T) {
	result := Solve("test.txt", true)
	if result != 46 {
		t.Errorf("Part 1 returned %d, expected %d", result, 46)
	}
}

func TestPart2(t *testing.T) {
	result := Solve("test.txt", false)
	if result != 51 {
		t.Errorf("Part 2 returned %d, expected %d", result, 51)
	}
}
