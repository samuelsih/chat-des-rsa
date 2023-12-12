package des

import (
	"log"
	"strconv"
	"strings"

	"github.com/samuelsih/chat-encrypt/des/table"
)

// 4 bit * 8 iteration = 32 bit
func substitute(s []string) []string {
	dividedBlocks := [][]string{s[:6], s[6:12], s[12:18], s[18:24], s[24:30], s[30:36], s[36:42], s[42:]}

	var resultString string
	var substitutionResult int

	for i, dividedBlock := range dividedBlocks {
		dividedBlock := dividedBlock

		x := getMiddleBits(dividedBlock)       // 2,3,4,5
		y := getFirstAndLastBits(dividedBlock) // 1,6

		switch i {
		case 0:
			substitutionResult = table.Subs1(x, y)
		case 1:
			substitutionResult = table.Subs2(x, y)
		case 2:
			substitutionResult = table.Subs3(x, y)
		case 3:
			substitutionResult = table.Subs4(x, y)
		case 4:
			substitutionResult = table.Subs5(x, y)
		case 5:
			substitutionResult = table.Subs6(x, y)
		case 6:
			substitutionResult = table.Subs7(x, y)
		case 7:
			substitutionResult = table.Subs8(x, y)
		}

		resultString = addPaddingZeros(substitutionResult, resultString)
	}

	return strings.Split(resultString, "")
}

// 1 -> 0001
func addPaddingZeros(substitutionResult int, resultString string) string {
	switch {
	case substitutionResult <= 1:
		return resultString + "000" + strconv.FormatUint(uint64(substitutionResult), 2)
	case substitutionResult <= 3:
		return resultString + "00" + strconv.FormatUint(uint64(substitutionResult), 2)
	case substitutionResult <= 7:
		return resultString + "0" + strconv.FormatUint(uint64(substitutionResult), 2)
	default:
		return resultString + strconv.FormatUint(uint64(substitutionResult), 2)
	}
}

func getMiddleBits(dividedBlock []string) uint64 {
	result, err := strconv.ParseUint(strings.Join(dividedBlock[1:5], ""), 2, 0)
	if err != nil {
		log.Fatalf("getMiddleBits: %v", err)
	}

	return result
}

func getFirstAndLastBits(dividedBlock []string) uint64 {
	result, err := strconv.ParseUint(strings.Join([]string{dividedBlock[0], dividedBlock[5]}, ""), 2, 0)
	if err != nil {
		log.Fatalf("getFirstAndLastBits: %v", err)
	}

	return result
}

func enciphering(ip []string, keys []string, decrypt bool) ([]string, []string) {
	leftBlock := ip[:32]
	rightBlock := ip[32:]

	for i := 0; i < 16; i++ {
		expandedRightBlock := table.ExpansionPermutationTable(rightBlock)
		var keyInt uint64
		var err error

		if decrypt == false {
			keyInt, err = strconv.ParseUint(keys[i], 2, 0)
		} else if decrypt == true {
			keyInt, err = strconv.ParseUint(keys[15-i], 2, 0)
		}

		if err != nil {
			log.Fatalf("enciphering 1: %v", err)
		}

		xoredKeyWithRightBlock := xorRightBlockWithKey(expandedRightBlock, keyInt)
		substitutedAndPermutePRightBlock := substituteRightBlockAndPermuteP(xoredKeyWithRightBlock)
		xoredLeftAndRight := xorRightBlockWithLeftBlock(leftBlock, substitutedAndPermutePRightBlock)

		leftBlock = rightBlock
		rightBlock = xoredLeftAndRight
	}

	return leftBlock, rightBlock
}

func xorRightBlockWithKey(block []string, key uint64) []string {
	blockInt, err := strconv.ParseUint(strings.Join(block, ""), 2, 0)
	if err != nil {
		log.Fatalf("xorRightBlockWithKey.blockInt: %v", err)
	}

	xored := blockInt ^ key
	xoredStr := strconv.FormatUint(xored, 2)
	for len(xoredStr) < 48 {
		xoredStr = "0" + xoredStr
	}

	xoredBlock := strings.Split(xoredStr, "")
	return xoredBlock
}

func substituteRightBlockAndPermuteP(rightBlock []string) []string {
	substitutedBlock := substitute(rightBlock)
	substitutedBlock = table.PermutationPTable(substitutedBlock)

	return substitutedBlock
}

func xorRightBlockWithLeftBlock(left, right []string) []string {
	rightBlockInt, err := strconv.ParseUint(strings.Join(right, ""), 2, 0)
	if err != nil {
		log.Fatalf("xorRightBlockWithLeftBlock.rightBlockInt: %v", err)
	}

	leftBlockInt, err := strconv.ParseUint(strings.Join(left, ""), 2, 0)
	if err != nil {
		log.Fatalf("xorRightBlockWithLeftBlock.leftBlockInt: %v", err)
	}

	xored := rightBlockInt ^ leftBlockInt
	xoredStr := strconv.FormatUint(xored, 2)
	for len(xoredStr) < 32 {
		xoredStr = "0" + xoredStr
	}

	xoredBlock := strings.Split(xoredStr, "")
	return xoredBlock
}