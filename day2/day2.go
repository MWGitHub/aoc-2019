package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	sum      = 1
	multiply = 2
	halt     = 99
)

const (
	min = 0
	max = 99
)

// IntCoder calculates the IntCode for codes and returns a new array of codes.
func IntCoder(codes []int, noun int, verb int) ([]int, error) {
	intCodes := make([]int, len(codes))
	copy(intCodes, codes)
	intCodes[1] = noun
	intCodes[2] = verb

	for i := 0; i < len(intCodes); i += 4 {
		op := intCodes[i]
		if op == halt {
			break
		}
		// range check
		length := len(intCodes)
		if length-i < 4 {
			return nil, errors.New("indices out of bounds")
		}
		pos1 := intCodes[i+1]
		pos2 := intCodes[i+2]
		outPos := intCodes[i+3]
		if pos1 >= length || pos2 >= length || outPos >= length {
			return nil, errors.New("positions out of bounds")
		}

		x := intCodes[pos1]
		y := intCodes[pos2]
		if op == sum {
			intCodes[outPos] = x + y
		} else if op == multiply {
			intCodes[outPos] = x * y
		} else {
			return nil, errors.New("invalid opcode")
		}
	}

	return intCodes, nil
}

// FindNounAndVerb finds the noun and verb that will result in the desired output.
// Output returns -1 and the error when an error occurs.
// Brute force method with O(n)
// More specific: O(c*m*n) with m maximum as max(100,n)^2 and c for array intialization and copying
func FindNounAndVerb(codes []int, desiredOutput int) (int, error) {
	for noun := min; noun <= max; noun++ {
		for verb := min; verb <= max; verb++ {
			result, err := IntCoder(codes, noun, verb)
			if err != nil {
				continue
			}

			if result[0] == desiredOutput {
				return 100*noun + verb, nil
			}
		}
	}
	return -1, errors.New("no noun and verb combination possible")
}

func main() {
	argInput := os.Args[1]

	input := strings.Split(argInput, ",")
	codes := make([]int, len(input))
	for i, v := range input {
		code, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		codes[i] = code
	}
	// starting with 1202 program alarm state
	result, err := IntCoder(codes, 12, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result[0])
	// finding a noun and verb with an output of 19690720
	const desiredOutput = 19690720
	nounVerb, err := FindNounAndVerb(codes, desiredOutput)
	fmt.Println(nounVerb)
}
