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
