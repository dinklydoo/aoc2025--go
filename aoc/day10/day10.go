package day10

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type factory struct {
	config  []bool
	buttons [][]int
	joltage []int
}

func formatInput(input string) []factory {
	lines := strings.Split(input, "\n")
	factories := make([]factory, len(lines))

	for i, line := range lines {
		tokens := strings.Split(line, " ")
		factory := &factories[i]
		for _, token := range tokens {
			delim := token[1 : len(token)-1]

			switch token[0] {
			case '[':
				for _, machine := range delim {
					if machine == '#' {
						factory.config = append(factory.config, true)
					} else {
						factory.config = append(factory.config, false)
					}
				}
			case '(':
				args := strings.Split(delim, ",")
				var temp []int
				for _, arg := range args {
					val, _ := strconv.Atoi(arg)
					temp = append(temp, val)
				}
				factory.buttons = append(factory.buttons, temp)
			case '{':
				jolts := strings.Split(delim, ",")
				for _, jolt := range jolts {
					val, _ := strconv.Atoi(jolt)
					factory.joltage = append(factory.joltage, val)
				}
			}
		}
	}
	return factories
}

func Part1(input string) int {
	factories := formatInput(input)
	// this is NP-hard (SAT-like) so brute force with power set, our input seems limited by like 10 button configurations so pre chill
	res := 0
	for i := range factories {
		factory := &factories[i]

		buttons := len(factory.buttons)
		machines := len(factory.config)

		temp := math.MaxInt

		bitmask := (1 << buttons) - 1
		for pset := 0; pset <= bitmask; pset++ {
			toggle := make([]bool, machines)

			count := 0
			sel := 1
			for i := 0; i < buttons; i++ {
				if pset&sel > 0 {
					count++
					for _, sw := range factory.buttons[i] {
						toggle[sw] = toggle[sw] != true
					}
				}
				sel <<= 1
			}

			can := true
			for i := 0; i < machines; i++ {
				if toggle[i] != factory.config[i] {
					can = false
					break
				}
			}
			if can {
				temp = min(temp, count)
			}
		}
		res += temp
	}
	return res
}

func Part2(input string) int {
	factories := formatInput(input)
	res := 0

	for _, factory := range factories {
		n := len(factory.joltage)
		m := len(factory.buttons) + 1

		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		for i, button := range factory.buttons {
			for _, b := range button {
				matrix[b][i] = 1
			}
		}
		for i, jolt := range factory.joltage {
			matrix[i][m-1] = jolt
		}

		matrix = Reduce(matrix) // int row form of linear system

		temp := Solve(matrix)
		fmt.Println(temp)
		res += temp
	}
	return res
}
