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

// IntCoder calculates the IntCode for codes.
// Out of bounds are not handled.
func IntCoder(codes []int) []int {
	intCodes := make([]int, len(codes))
	copy(intCodes, codes)

	for i := 0; i < len(intCodes); i += 4 {
		op := intCodes[i]
		if op == halt {
			break
		}

		x := intCodes[intCodes[i+1]]
		y := intCodes[intCodes[i+2]]
		outPos := intCodes[i+3]
		if op == sum {
			intCodes[outPos] = x + y
		} else if op == multiply {
			intCodes[outPos] = x * y
		} else {
			panic(errors.New("invalid opcode"))
		}
	}

	return intCodes
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
	// start with 1202 program alarm state
	codes[1] = 12
	codes[2] = 2

	fmt.Println(IntCoder(codes)[0])
}
