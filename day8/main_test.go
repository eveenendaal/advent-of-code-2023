package main

import "testing"

func TestPart1(t *testing.T) {
	test1 := Part1("test1.txt")
	if test1 != 2 {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", test1, 2)
	}

	test2 := Part1("test2.txt")
	if test2 != 6 {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", test2, 6)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Write test cases for Function2
}
