package main

import (
	"errors"
	"fmt"
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
	fmt.Println("foo")
}
