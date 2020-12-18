package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func compute_program(program []int) []int {
	memory := make([]int, len(program), cap(program))
	copy(memory, program)
	program_counter := 0

	for memory[program_counter] != -1 {
		if memory[program_counter] == 1 {
			read_location_a := memory[program_counter+1]
			read_location_b := memory[program_counter+2]
			write_location := memory[program_counter+3]

			memory[write_location] = memory[read_location_a] + memory[read_location_b]
		} else if memory[program_counter] == 2 {
			read_location_a := memory[program_counter+1]
			read_location_b := memory[program_counter+2]
			write_location := memory[program_counter+3]

			memory[write_location] = memory[read_location_a] * memory[read_location_b]
		} else if memory[program_counter] == 99 {
			// fmt.Println("Instruction 99")
			break
		} else if program_counter > len(memory) && program_counter != -1 {
			fmt.Println("Buffer overflow! Cannot read this location. Does not exist!")
		} else {
			fmt.Println("Unrecognized instruction", memory[program_counter])
			program_counter = -1
		}
		program_counter = program_counter + 4
	}

	return memory
}

func generate_symbols() []int {
	s := make([]int, 100)
	for i := range s {
		s[i] = i
		i = i + 1
	}
	return s
}

func main() {
	f, err := os.Open("day02-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := csv.NewReader(f)

	program_csv, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(program_csv)

	original_program := make([]int, cap(program_csv[0]))
	program := make([]int, cap(program_csv[0]))

	for i, v := range program_csv[0] {
		original_program[i], err = strconv.Atoi(v)

		if err != nil {
			log.Fatal(err)
		}
	}

	copy(program, original_program)

	program[1] = 12
	program[2] = 2
	result := compute_program(program)
	fmt.Println("Day 02 Part 1:", result[0])

	// Part 2
	copy(program, original_program)

	nouns := generate_symbols()
	verbs := generate_symbols()
	combinations := len(nouns) * len(verbs)

	// fmt.Println(nouns)
	// fmt.Println(verbs)
	fmt.Println(combinations)

	// I can do better here, but **shrug**
	answer := 0
	expected_answer := 19690720
	final_noun := 0
	final_verb := 0

	for _, noun := range nouns {
		copy(program, original_program)
		answer = 0
		for _, verb := range verbs {
			// fmt.Println(noun, verb)
			// fmt.Println(noun, verb)
			program[1] = noun
			program[2] = verb
			result := compute_program(program)
			answer = result[0]
			final_noun = noun
			final_verb = verb
			if answer == expected_answer {
				fmt.Println("I have found an answer!")
				fmt.Println(final_noun, final_verb)
				fmt.Println("Day 02 Part 2:", (100*final_noun)+final_verb)
			}
		}
	}
}
