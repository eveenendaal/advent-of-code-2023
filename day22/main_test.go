package day22

import "testing"

func TestPart1(t *testing.T) {
	want := 5
	got := Part1("test.txt")
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}
