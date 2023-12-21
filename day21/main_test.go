package day21

import "testing"

func TestPart1Test(t *testing.T) {
	result := Part1("test.txt", 6)
	if result != 16 {
		t.Errorf("Expected result to be 16, got %d", result)
	}
}
