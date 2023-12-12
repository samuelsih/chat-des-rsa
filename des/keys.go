package des

import (
	"strings"

	"github.com/samuelsih/chat-encrypt/des/table"
)

var leftShiftIndex = []int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

// frontliner to last index
func circularLeftShift(input []string, times int) []string {
	if times == 1 {
		first := input[0]

		for i := range input {
			if i == len(input)-1 {
				input[i] = first
			} else {
				input[i] = input[i+1]
			}
		}

		return input
	}

	first := input[0]
	second := input[1]

	for i := range input {
		switch {
		case i == len(input)-2:
			input[i] = first
		case i == len(input)-1:
			input[i] = second
		default:
			input[i] = input[i+2]
		}
	}

	return input
}

// generate 16 subkeys, each 48 bits
func generateKeys(str string) []string {
	var keys []string

	initialKey := StrToBinary(str)
	initialKeyBlock := strings.Split(initialKey, "")

	leftBlock, rightBlock := table.PermuteChoice1(initialKeyBlock)

	for i := 0; i < 16; i++ {
		leftBlock = circularLeftShift(leftBlock, leftShiftIndex[i])
		rightBlock = circularLeftShift(rightBlock, leftShiftIndex[i])

		concatenatedBlock := append(leftBlock, rightBlock...)

		pc2 := table.PermuteChoice2(concatenatedBlock)

		keys = append(keys, strings.Join(pc2, ""))
	}

	return keys
}