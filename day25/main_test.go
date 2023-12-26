package day25

import "testing"

func TestPart1Test(t *testing.T) {
	want := 54
	if got := Part1("test.txt"); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart1Solution(t *testing.T) {
	want := 562978
	if got := Part1("input.txt"); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}
