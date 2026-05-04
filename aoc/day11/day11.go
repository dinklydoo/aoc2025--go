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

func dfs2(conn map[string][]string, current string, fft bool, dac bool) int {
	if current == "out" {
		if fft && dac {
			return 1
		}
		return 0
	}

	isFFT := current == "fft"
	isDAC := current == "dac"

	res := 0
	for _, next := range conn[current] {
		res += dfs2(conn, next, fft || isFFT, dac || isDAC)
	}
	return res
}

// TODO : too slow, need to do some topo sort shenanigans, issue is nodes are restricted to be funneled bw fft and dta
/* Traversal paths looks like this
	out
//// \\\ \\
	...
\\ \ / \ // /
	fft
//// \\\ \\
	...
\\	\\///
	dta
//// \\\ \\
	...
	\ //
	out
*/

func Part2(input string) int {
	conn := formatInput(input)
	return dfs2(conn, "svr", false, false)
}
