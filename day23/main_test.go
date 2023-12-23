package day23

import (
	"testing"
)

func TestPart1Test(t *testing.T) {
	expected := 74
	actual := Part1("test.txt")
	if actual != expected {
		t.Errorf("Part1() = %d, expected %d", actual, expected)
	}
}

func TestPart1Solution(t *testing.T) {
	expected := 0
	actual := Part1("input.txt")
	if actual != expected {
		t.Errorf("Part1() = %d, expected %d", actual, expected)
	}
}
