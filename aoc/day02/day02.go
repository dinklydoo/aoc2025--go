package day02

import (
	"aoc2025--go/aoc/utils"
	"strconv"
	"strings"
)

// format the input file to [2]int ranges
func formatInput(input string) [][2]int {
	split := strings.Split(input, ",")
	var ranges = make([][2]int, len(split))
	for i, ran := range split {
		pair := strings.Split(ran, "-")

		l, _ := strconv.Atoi(pair[0])
		r, _ := strconv.Atoi(pair[1])

		ranges[i] = [2]int{l, r}
	}
	return ranges
}

// count number of digits in a number
func digits(n int) int {
	length := 0
	for n > 0 {
		length++
		n /= 10
	}
	return length
}

// check if a number can be split into m repeating fragments
func checkSplit(n int, m int) bool {
	length := digits(n)
	utils.Assert(length%m == 0, "number must be splittable into equal fragments")

	frag := -1
	for i := 0; i < m; i++ {
		base := 1
		curr := 0

		for j := 0; j < length/m; j++ {
			curr += n % 10 * base
			base *= 10
			n /= 10
		}

		if frag < 0 {
			frag = curr
		}

		if curr != frag {
			return false
		}
	}
	return true
}

func Part1(input string) int {
	ranges := formatInput(input)

	res := 0
	for _, ran := range ranges {
		for i := ran[0]; i <= ran[1]; i++ {
			if digits(i)&1 == 1 {
				continue
			}
			if checkSplit(i, 2) {
				res += i
			}
		}
	}
	return res
}

func Part2(input string) int {
	primes := utils.Primes(20)
	factors := make([][]int, 20)
	for i := range factors {
		factors[i] = utils.Factorize(i, primes)
	}

	ranges := formatInput(input)
	res := 0
	for _, ran := range ranges {
		for i := ran[0]; i <= ran[1]; i++ {
			length := digits(i)

			for _, factor := range factors[length] {
				if checkSplit(i, factor) {
					res += i
					break
				}
			}

		}
	}
	return res
}
