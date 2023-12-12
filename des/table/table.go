package table

func PermuteChoice1(input []string) ([]string, []string) {
	return []string{
			input[56], input[48], input[40], input[32], input[24], input[16], input[8],
			input[0], input[57], input[49], input[41], input[33], input[25], input[17],
			input[9], input[1], input[58], input[50], input[42], input[34], input[26],
			input[18], input[10], input[2], input[59], input[51], input[43], input[35]},
		[]string{
			input[62], input[54], input[46], input[38], input[30], input[22], input[14],
			input[6], input[61], input[53], input[45], input[37], input[29], input[21],
			input[13], input[5], input[60], input[52], input[44], input[36], input[28],
			input[20], input[12], input[4], input[27], input[19], input[11], input[3]}
}

func InitialPermutation(input []string) []string {
	return []string{
		input[57], input[49], input[41], input[33], input[25], input[17], input[9], input[1],
		input[59], input[51], input[43], input[35], input[27], input[19], input[11], input[3],
		input[61], input[53], input[45], input[37], input[29], input[21], input[13], input[5],
		input[63], input[55], input[47], input[39], input[31], input[23], input[15], input[7],
		input[56], input[48], input[40], input[32], input[24], input[16], input[8], input[0],
		input[58], input[50], input[42], input[34], input[26], input[18], input[10], input[2],
		input[60], input[52], input[44], input[36], input[28], input[20], input[12], input[4],
		input[62], input[54], input[46], input[38], input[30], input[22], input[14], input[6],
	}
}

func InverseInitialPermutation(input []string) []string {
	return []string{
		input[39], input[7], input[47], input[15], input[55], input[23], input[63], input[31],
		input[38], input[6], input[46], input[14], input[54], input[22], input[62], input[30],
		input[37], input[5], input[45], input[13], input[53], input[21], input[61], input[29],
		input[36], input[4], input[44], input[12], input[52], input[20], input[60], input[28],
		input[35], input[3], input[43], input[11], input[51], input[19], input[59], input[27],
		input[34], input[2], input[42], input[10], input[50], input[18], input[58], input[26],
		input[33], input[1], input[41], input[9], input[49], input[17], input[57], input[25],
		input[32], input[0], input[40], input[8], input[48], input[16], input[56], input[24],
	}
}

func PermuteChoice2(input []string) []string {
	return []string{
		input[13], input[16], input[10], input[23], input[0], input[4], input[2], input[27],
		input[14], input[5], input[20], input[9], input[22], input[18], input[11], input[3],
		input[25], input[7], input[15], input[6], input[26], input[19], input[12], input[1],
		input[40], input[51], input[30], input[36], input[46], input[54], input[29], input[39],
		input[50], input[44], input[32], input[47], input[43], input[48], input[38], input[55],
		input[33], input[52], input[45], input[41], input[49], input[35], input[28], input[31],
	}
}

func ExpansionPermutationTable(input []string) []string {
	return []string{
		input[31], input[0], input[1], input[2], input[3], input[4], input[3], input[4],
		input[5], input[6], input[7], input[8], input[7], input[8], input[9], input[10],
		input[11], input[12], input[11], input[12], input[13], input[14], input[15], input[16],
		input[15], input[16], input[17], input[18], input[19], input[20], input[19], input[20],
		input[21], input[22], input[23], input[24], input[23], input[24], input[25], input[26],
		input[27], input[28], input[27], input[28], input[29], input[30], input[31], input[0],
	}
}

func PermutationPTable(input []string) []string {
	return []string{
		input[15], input[6], input[19], input[20], input[28], input[11], input[27], input[16],
		input[0], input[14], input[22], input[25], input[4], input[17], input[30], input[9],
		input[1], input[7], input[23], input[13], input[31], input[26], input[2], input[8],
		input[18], input[12], input[29], input[5], input[21], input[10], input[3], input[24],
	}
}