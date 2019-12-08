// Calculates fuel requirements for modules.

package main

import (
	"fmt"
	"os"
	"strconv"
)

// CalculateFuelRequirements returns the fuel needed for a given mass.
func CalculateFuelRequirements(mass int) int {
	divider := 3
	tolerance := 2

	if fuel := (mass / divider) - tolerance; fuel > 0 {
		return fuel + CalculateFuelRequirements(fuel)
	}

	return 0
}

func main() {
	masses := os.Args[1:]

	fuel := 0
	for _, input := range masses {
		mass, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		fuel += CalculateFuelRequirements(mass)
	}

	fmt.Println(fuel)
}
