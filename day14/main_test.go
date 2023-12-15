package main

import "testing"

func TestFunction1(t *testing.T) {
	result := Solve("test.txt")
	if result != 136 {
		t.Errorf("Expected 136, got %d", result)
	}
}

func TestFunction2(t *testing.T) {
	// TODO: Add your test logic here
}
