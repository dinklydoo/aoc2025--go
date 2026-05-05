package day11

import (
	"fmt"
	"strings"
)

func formatInput(input string) map[string][]string {
	lines := strings.Split(input, "\n")

	conn := make(map[string][]string)
	for _, line := range lines {
		temp := strings.Split(line, ":")
		src := temp[0]

		temp = strings.Fields(temp[1])
		for _, t := range temp {
			conn[src] = append(conn[src], t)
		}
	}
	return conn
}

func dfs(conn map[string][]string, current string) int {
	fmt.Println(current)
	if current == "out" {
		return 1
	}
	res := 0
	for _, next := range conn[current] {
		res += dfs(conn, next)
	}
	return res
}

func Part1(input string) int {
	conn := formatInput(input)
	return dfs(conn, "you")
}

type state struct {
	node string
	fft  bool
	dta  bool
}

func dfs2(conn map[string][]string, cache map[state]int, current string, fft bool, dac bool) int {
	if current == "out" {
		if fft && dac {
			return 1
		}
		return 0
	}

	isFFT := current == "fft"
	isDAC := current == "dac"

	state := state{node: current, fft: fft, dta: dac}
	if val, ok := cache[state]; ok {
		return val
	}

	res := 0
	for _, next := range conn[current] {
		res += dfs2(conn, cache, next, fft || isFFT, dac || isDAC)
	}
	cache[state] = res
	return res
}

func Part2(input string) int {
	conn := formatInput(input)
	cache := make(map[state]int)

	return dfs2(conn, cache, "svr", false, false)
}
