package day01

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	ins := strings.Split(input, "\n")

	res := 0
	pos := 50
	for _, in := range ins {

		rot, _ := strconv.Atoi(in[1:])

		if in[0] == 'R' {
			pos = (pos + rot) % 100
		} else {
			pos = (pos - rot + 100) % 100
		}

		if pos == 0 {
			res++
		}
	}
	return res
}

func Part2(input string) int {
	ins := strings.Split(input, "\n")

	res := 0
	pos := 50
	for _, in := range ins {
		rot, _ := strconv.Atoi(in[1:])

		res += rot / 100 // cycles will always click on each cycle

		rot %= 100
		if in[0] == 'R' {
			if pos+rot >= 100 {
				res++
			}
			pos = (pos + rot) % 100
		} else {
			// starting at position 0 -> no way to click again
			if pos > 0 && pos-rot <= 0 {
				res++
			}
			pos = (pos - rot + 100) % 100
		}
	}
	return res
}
