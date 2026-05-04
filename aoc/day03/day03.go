package day03

import (
	"aoc2025--go/aoc/utils"
	"strings"
)

func formatInput(input string) []string {
	split := strings.Split(input, "\n")
	return split
}

func Part1(input string) int {
	banks := formatInput(input)

	res := 0
	for _, bank := range banks {
		temp := 0
		small := -1

		for i := len(bank) - 1; i >= 0; i-- {
			curr := int(bank[i] - '0')
			if small < 0 {
				small = curr
			} else {
				temp = max(temp, 10*curr+small)
			}
			small = max(small, curr)
		}
		res += temp
	}
	return res
}

func Part2(input string) int {
	banks := formatInput(input)

	res := 0
	for _, bank := range banks {
		temp := make([]int, 13)
		for i := 1; i <= 12; i++ {
			temp[i] = -1
		}

		for i := len(bank) - 1; i >= 0; i-- {
			curr := int(bank[i] - '0')

			base := utils.Pow(10, 11)
			for j := 12; j >= 1; j-- {
				if temp[j-1] >= 0 {
					temp[j] = max(temp[j], base*curr+temp[j-1])
				}
				base /= 10
			}
		}
		res += temp[12]
	}
	return res
}
