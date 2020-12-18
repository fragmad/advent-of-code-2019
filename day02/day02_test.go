package main

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestDay01Part1Examples(t *testing.T) {
	example1_program := []int{1, 0, 0, 0, 99}
	example1_answer := compute_program(example1_program)
	example1_expected := []int{2, 0, 0, 0, 99}

	if !Equal(example1_answer, example1_expected) {
		t.Errorf("Program did not end in expected state, got: %d, expected: %d", example1_answer, example1_expected)
	}

	example2_program := []int{2, 3, 0, 3, 99}
	example2_answer := compute_program(example2_program)
	example2_expected := []int{2, 3, 0, 6, 99}

	if !Equal(example2_answer, example2_expected) {
		t.Errorf("Program did not end in expected state, got: %d, expected: %d", example2_answer, example2_expected)
	}

	example3_program := []int{2, 4, 4, 5, 99, 0}
	example3_answer := compute_program(example3_program)
	example3_expected := []int{2, 4, 4, 5, 99, 9801}

	if !Equal(example3_answer, example3_expected) {
		t.Errorf("Program did not end in expected state, got: %d, expected: %d", example3_answer, example3_expected)
	}

	example4_program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	example4_answer := compute_program(example4_program)
	example4_expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

	if !Equal(example4_answer, example4_expected) {
		t.Errorf("Program did not end in expected state, got: %d, expected: %d", example4_answer, example4_expected)
	}
}
