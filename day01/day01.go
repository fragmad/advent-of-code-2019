package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*Fuel required to launch a given module is based on its mass.
Specifically, to find the fuel required for a module, take its mass, divide
by three, round down, and subtract 2.*/
func calculate_fuel(mass int) int {
	return (mass / 3) - 2
}

/* Fuel itself requires fuel just like a module - take its mass, divide by three,
round down, and subtract 2. However, that fuel also requires fuel, and that fuel
requires fuel, and so on. Any mass that would require negative fuel should instead
be treated as if it requires zero fuel; the remaining mass, if any, is instead handled
by wishing really hard, which has no mass and is outside the scope of this calculation.
*/
func calculate_complete_fuel(module_mass int) int {
	total_fuel := calculate_fuel(module_mass)
	mass_of_fuel := calculate_fuel(module_mass)

	for mass_of_fuel > 0 {
		mass_of_fuel = calculate_fuel(mass_of_fuel)

		if mass_of_fuel < 0 {
			mass_of_fuel = 0
		}

		// fmt.Println("Mass of fuel", mass_of_fuel)
		total_fuel = total_fuel + mass_of_fuel
	}

	return total_fuel
}

func sum_fuel_calculations(calculations []int) int {
	total := 0
	for _, c := range calculations {
		total = total + c
	}
	return total
}

func main() {
	f, err := os.Open("day01-input.txt")
	fuel_calculations_part1 := make([]int, 100)
	fuel_calculations_part2 := make([]int, 100)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		i, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		fuel_calculations_part1 = append(fuel_calculations_part1, calculate_fuel(i))
		fuel_calculations_part2 = append(fuel_calculations_part2, calculate_complete_fuel(i))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", sum_fuel_calculations(fuel_calculations_part1))
	fmt.Println("Part 2:", sum_fuel_calculations(fuel_calculations_part2))
}
