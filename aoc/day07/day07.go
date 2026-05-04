package day07

import (
	"strings"
)

func formatInput(input string) [][]byte {
	lines := strings.Split(input, "\n")
	n := len(lines)

	bytes := make([][]byte, n)

	for i, line := range lines {
		bytes[i] = []byte(line)
	}
	return bytes
}

func Part1(input string) int {
	bytes := formatInput(input)
	n := len(bytes)
	m := len(bytes[0])

	beams := make([][]bool, n)
	for i := 0; i < n; i++ {
		beams[i] = make([]bool, m)
	}

	res := 0
	for i := range bytes {
		for j, b := range bytes[i] {

			if i == 0 {
				beams[i][j] = b == 'S'
				continue
			}

			if bytes[i-1][j] == '^' {
				if beams[i-1][j] == true {
					res++
					if j > 0 {
						beams[i][j-1] = true
					}
					if j < m-1 {
						beams[i][j+1] = true
					}
				}
			} else {
				beams[i][j] = beams[i][j] || beams[i-1][j]
			}
		}
	}
	return res
}

func Part2(input string) int {
	bytes := formatInput(input)
	n := len(bytes)
	m := len(bytes[0])

	beams := make([][]int, n)
	for i := 0; i < n; i++ {
		beams[i] = make([]int, m)
	}

	res := 0
	for i := range bytes {
		for j, b := range bytes[i] {

			if i == 0 {
				if b == 'S' {
					beams[i][j] = 1
				}
				continue
			}

			if bytes[i-1][j] == '^' {
				if beams[i-1][j] > 0 {

					if j > 0 {
						beams[i][j-1] += beams[i-1][j]
					}
					if j < m-1 {
						beams[i][j+1] += beams[i-1][j]
					}
				}
			} else {
				beams[i][j] += beams[i-1][j]
			}
		}
	}

	for _, k := range beams[n-1] {
		res += k
	}

	return res
}
