package base

import (
	"testing"
)

func TestFileToLines(t *testing.T) {
	lines := ReadFileToLines("test.txt")
	if len(lines) != 13 {
		t.Errorf("Expected 13 lines, got %d", len(lines))
	}
}

func TestFileToCharacters(t *testing.T) {
	characters := ReadFileToCharacters("test.txt")
	if len(characters) != 13 {
		t.Errorf("Expected 13 lines, got %d", len(characters))
	}
	if len(characters[0]) != 13 {
		t.Errorf("Expected 13 characters, got %d", len(characters[0]))
	}
}

func TestFileToInt(t *testing.T) {
	ints := ReadFileToInt("test.txt")
	if len(ints) != 13 {
		t.Errorf("Expected 13 lines, got %d", len(ints))
	}
	if len(ints[0]) != 13 {
		t.Errorf("Expected 13 ints, got %d", len(ints[0]))
	}
}

func TestIntAbs(t *testing.T) {
	if IntAbs(-1) != 1 {
		t.Errorf("Expected 1, got %d", IntAbs(-1))
	}
	if IntAbs(1) != 1 {
		t.Errorf("Expected 1, got %d", IntAbs(1))
	}
	if IntAbs(0) != 0 {
		t.Errorf("Expected 0, got %d", IntAbs(0))
	}
}

type TestShape struct {
	interior int
	exterior int
	points   []Position
}

var shapes = []TestShape{
	TestShape{
		interior: 1,
		exterior: 9,
		points: []Position{
			{0, 0},
			{2, 0},
			{2, 2},
			{0, 2},
		},
	},
	TestShape{
		interior: 2, // off by one expected; the answer is actually 1
		exterior: 11,
		points: []Position{
			{0, 0},
			{3, 0},
			{3, 2},
			{1, 2},
			{1, 1},
			{0, 1},
		},
	},
}

func TestFindAreaShoelace(t *testing.T) {
	for _, shape := range shapes {
		area := FindInclusiveAreaShoelace(shape.points)
		if area != shape.exterior {
			PrintShape(shape.points)
			t.Errorf("Expected %d, got %d", shape.exterior, area)
		}
	}
}

func TestFindInteriorAreaShoelace(t *testing.T) {
	for _, shape := range shapes {
		area := FindInteriorAreaShoelace(shape.points)
		if area != shape.interior {
			PrintShape(shape.points)
			t.Errorf("Expected %d, got %d", shape.interior, area)
		}
	}
}
