package day04

import "strings"

func formatInput(input string) [][]byte {
	split := strings.Split(input, "\n")

	n := len(split)

	bytes := make([][]byte, n)
	for i := range bytes {
		bytes[i] = []byte(split[i])
	}
	return bytes
}

func sweepRolls(rolls [][]byte) (int, [][]byte) {
	n := len(rolls)
	m := len(rolls[0])
	nextRolls := make([][]byte, n)
	for i := range nextRolls {
		nextRolls[i] = make([]byte, m)
	}

	res := 0
	for i := range rolls {
		for j := range rolls[i] {
			nextRolls[i][j] = rolls[i][j]

			if rolls[i][j] != '@' {
				continue
			}

			adj := 0
			if i > 0 {
				if rolls[i-1][j] == '@' {
					adj++
				}
				if j > 0 && rolls[i-1][j-1] == '@' {
					adj++
				}
			}
			if i < n-1 {
				if rolls[i+1][j] == '@' {
					adj++
				}
				if j < m-1 && rolls[i+1][j+1] == '@' {
					adj++
				}
			}
			if j > 0 {
				if rolls[i][j-1] == '@' {
					adj++
				}
				if i < n-1 && rolls[i+1][j-1] == '@' {
					adj++
				}
			}
			if j < m-1 {
				if rolls[i][j+1] == '@' {
					adj++
				}
				if i > 0 && rolls[i-1][j+1] == '@' {
					adj++
				}
			}

			// increment and clear
			if adj < 4 {
				res++
				nextRolls[i][j] = '.'
			}
		}
	}
	return res, nextRolls
}

func Part1(input string) int {
	rolls := formatInput(input)

	res, _ := sweepRolls(rolls)
	return res
}

func Part2(input string) int {
	rolls := formatInput(input)

	res := 0
	temp, rolls := sweepRolls(rolls)

	res += temp
	for temp > 0 {
		temp, rolls = sweepRolls(rolls)
		res += temp
	}
	return res
}
