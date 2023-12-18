package main

import "testing"

func TestFunction1(t *testing.T) {
	result := Solve("test.txt")
	if result != 62 {
		t.Errorf("Expected 62, got %d", result)
	}
}

func TestFunction2(t *testing.T) {
	// TODO: Implement test
}

func TestFunction3(t *testing.T) {
	// TODO: Implement test
}

func TestFunction4(t *testing.T) {
	// TODO: Implement test
}
