package day22

import "testing"

func TestPart1Test(t *testing.T) {
	want := 5
	got := Part1("test.txt")
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart1Solution(t *testing.T) {
	want := 499
	got := Part1("input.txt")
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2Test(t *testing.T) {
	want := 7
	got := Part2("test.txt")
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func TestPart2Solution(t *testing.T) {
	want := 95059
	got := Part2("input.txt")
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}
