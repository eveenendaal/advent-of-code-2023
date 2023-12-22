package day21

import "testing"

func TestPart1Test(t *testing.T) {
	result := Part1("test.txt", 6)
	if result != 16 {
		t.Errorf("Expected result to be 16, got %d", result)
	}
}

func TestPart1Solution(t *testing.T) {
	result := Part1("input.txt", 64)
	if result != 3746 {
		t.Errorf("Expected result to be 3746, got %d", result)
	}
}

func TestPart2Test(t *testing.T) {
	answers := map[int]int64{
		6:   16,
		10:  50,
		50:  1594,
		100: 6536,
		500: 167004,
		//1000: 668697,
		//5000: 16733044,
	}

	for steps, answer := range answers {
		result := Part1("test.txt", steps)
		if result != answer {
			t.Errorf("Expected result to be %d, got %d", answer, result)
		} else {
			t.Logf("Steps: %d, result: %d", steps, result)
		}
	}
}

func TestPart2Solution(t *testing.T) {
	result := Part2("input.txt", 26501365)
	if result != 623540829615589 {
		t.Errorf("Expected result to be 623540829615589, got %d", result)
	}
}
