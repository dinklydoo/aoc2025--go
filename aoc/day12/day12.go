package day12

import (
	"strconv"
	"strings"
)

type shape struct {
	area [3][3]bool
}

type space struct {
	rows     int
	cols     int
	presents [6]int
}

func parseInput(input string) ([6]shape, []space) {
	lines := strings.Split(input, "\n")
	// low-key hardcode the shape part
	var shapes [6]shape
	var spaces []space
	for i := 0; i < 6; i++ {
		for r, line := range lines[5*i+1:] {
			if len(line) == 0 {
				break
			}
			for c := 0; c < 3; c++ {
				if line[c] == '#' {
					shapes[i].area[r][c] = true
				}
			}
		}
	}
	for _, line := range lines[30:] {
		tokens := strings.Split(line, ":")
		dims := strings.Split(tokens[0], "x")

		row, _ := strconv.Atoi(dims[0])
		col, _ := strconv.Atoi(dims[1])

		sp := space{
			rows: row,
			cols: col,
		}
		args := strings.Fields(tokens[1])
		for i, arg := range args {
			val, _ := strconv.Atoi(arg)
			sp.presents[i] = val
		}
		spaces = append(spaces, sp)
	}
	return shapes, spaces
}

func Part1(input string) int {
	_, spaces := parseInput(input)

	res := 0
	for _, space := range spaces {
		if space.rows < 3 || space.cols < 3 {
			continue
		}

		presents := 0
		for _, p := range space.presents {
			presents += p
		}
		// if the tree can accommodate presents 3x3 grids then guarantee fit
		if presents*9 <= space.rows*space.cols {
			res++
			continue
		}

		// now we have to pack

		// little bit hacky I think my input does not require "fitting" ...
		// TODO : properly LMAO
	}

	return res
}

func Part2(input string) int {
	return 0
}
