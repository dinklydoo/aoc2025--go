package day06

import (
	"strconv"
	"strings"
)

type equation struct {
	args     []string
	length   int
	operator func(int, int) int
}

func formatInput1(input string) []equation {
	lines := strings.Split(input, "\n")
	var equations []equation

	constructed := false
	for _, line := range lines {
		tokens := strings.Fields(line)

		if !constructed {
			for _ = range tokens {
				equations = append(equations, equation{})
			}
			constructed = true
		}

		for i, token := range tokens {
			if token == "+" {
				equations[i].operator = func(a, b int) int { return a + b }
			} else if token == "*" {
				equations[i].operator = func(a, b int) int { return a * b }
			} else {
				equations[i].args = append(equations[i].args, token)
			}
		}
	}
	return equations
}

func formatInput2(input string) []equation {
	lines := strings.Split(input, "\n")
	var equations []equation

	ops := strings.Fields(lines[len(lines)-1])
	for i, op := range ops {
		equations = append(equations, equation{})

		if op == "+" {
			equations[i].operator = func(a, b int) int { return a + b }
		} else {
			equations[i].operator = func(a, b int) int { return a * b }
		}
	}

	// really disgusting string processing below
	pos := make([]int, len(lines)-1)
	for i := range equations {
		eq := &equations[i]

		length := 0
		for i := 0; i < len(lines)-1; i++ {
			p := pos[i]
			for p < len(lines[i]) && lines[i][p] != ' ' {
				p++
			}
			length = max(length, p-pos[i])
		}
		eq.length = length

		for i := 0; i < len(lines)-1; i++ {
			line := lines[i]
			p := pos[i]

			eq.args = append(eq.args, line[p:p+length])
			pos[i] += length + 1
		}
	}

	return equations
}

func Part1(input string) int {
	equations := formatInput1(input)
	res := 0

	for _, equation := range equations {
		temp, _ := strconv.Atoi(equation.args[0])

		for i := 1; i < len(equation.args); i++ {
			val, _ := strconv.Atoi(equation.args[i])
			temp = equation.operator(temp, val)
		}
		res += temp
	}

	return res
}

func Part2(input string) int {
	equations := formatInput2(input)
	res := 0

	for _, equation := range equations {

		cephargs := make([]int, equation.length)

		bases := make([]int, equation.length)
		for i := range bases {
			bases[i] = 1
		}

		for i := len(equation.args) - 1; i >= 0; i-- {
			arg := equation.args[i]

			for pos := equation.length - 1; pos >= 0; pos-- {
				if arg[pos] == ' ' {
					continue
				}
				cephargs[pos] += bases[pos] * int(arg[pos]-'0')
				bases[pos] *= 10
			}
		}

		temp := cephargs[0]
		for i, arg := range cephargs {
			if i == 0 {
				continue
			}
			temp = equation.operator(temp, arg)
		}
		res += temp
	}
	return res
}
