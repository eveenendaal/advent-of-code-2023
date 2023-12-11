package main

import "testing"

func TestPart1(t *testing.T) {
	result := Solve("test.txt", 2)
	if result != 374 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 374)
	}
}

func TestPart2(t *testing.T) {
	result := Solve("test.txt", 10)
	if result != 1030 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1030)
	}

	result = Solve("test.txt", 100)
	if result != 8410 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 8410)
	}
}
