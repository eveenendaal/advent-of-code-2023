package day25

import "testing"

func TestPart1Test(t *testing.T) {
	want := 0
	if got := Part1("test.txt"); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}
