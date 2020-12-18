package main

import "testing"

func TestDay03Part1Examples(t *testing.T) {
	example1_input := `R75,D30,R83,U83,L12,D49,R71,U7,L72
	U62,R66,U55,R34,D71,R55,D58,R83`

	example1_answer := compute_wire_cross_distance(example1_input)
	example1_expected_distance := 159

	if example1_answer != example1_expected_distance {
		t.Errorf("Program did not end in expected state, got: %d, expected: %d", example1_answer, example1_expected_distance)
	}
}
