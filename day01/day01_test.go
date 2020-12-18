package main

import "testing"

func Sum(a int, b int) int {
	return a + b
}
func TestDay01Part1Examples(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	// }

	example1_answer := calculate_fuel(12)
	example1_expected := 2

	if example1_answer != example1_expected {
		t.Errorf("Fuel calculation was incorrect, got: %d, expected: %d", example1_answer, example1_expected)
	}

	example2_answer := calculate_fuel(14)
	example2_expected := 2

	if example2_answer != example2_expected {
		t.Errorf("Fuel calculation was incorrect, got: %d, expected: %d", example2_answer, example2_expected)
	}

	example3_answer := calculate_fuel(1969)
	example3_expected := 654

	if example3_answer != example3_expected {
		t.Errorf("Fuel calculation was incorrect, got: %d, expected: %d", example2_answer, example2_expected)
	}

	example4_answer := calculate_fuel(100756)
	example4_expected := 33583
	if example4_answer != example4_expected {
		t.Errorf("Fuel calculation was incorrect, got: %d, expected: %d", example2_answer, example2_expected)
	}
}

func TestDay01Part2Examples(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	// }

	example1_answer := calculate_fuel(12)
	example1_expected := 2

	if example1_answer != example1_expected {
		t.Errorf("Fuel calculation was incorrect, got: %d, expected: %d", example1_answer, example1_expected)
	}

	example2_answer := calculate_complete_fuel(14)
	example2_expected := 2

	if example2_answer != example2_expected {
		t.Errorf("Fuel calculation was incorrect, got: %d, expected: %d", example2_answer, example2_expected)
	}

	example3_answer := calculate_complete_fuel(1969)
	example3_expected := 966

	if example3_answer != example3_expected {
		t.Errorf("Fuel calculation was incorrect, got: %d, expected: %d", example3_answer, example3_expected)
	}

	example4_answer := calculate_complete_fuel(100756)
	example4_expected := 50346
	if example4_answer != example4_expected {
		t.Errorf("Fuel calculation was incorrect, got: %d, expected: %d", example3_answer, example3_expected)
	}
}
