package day05

import (
	"slices"
	"strconv"
	"strings"
)

func formatInput(input string) ([][2]int, []int) {
	split := strings.Split(input, "\n")

	var ranges [][2]int
	var ingredients []int

	for _, line := range split {
		ran := strings.Split(line, "-")
		if len(ran) == 1 {
			ingredient, _ := strconv.Atoi(ran[0])
			ingredients = append(ingredients, ingredient)
		} else {
			l, _ := strconv.Atoi(ran[0])
			r, _ := strconv.Atoi(ran[1])

			ranges = append(ranges, [2]int{l, r})
		}
	}

	// sort slices non-dec start range order
	slices.SortFunc(ranges, func(a, b [2]int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	return ranges, ingredients
}

func isFresh(ranges [][2]int, id int) bool {
	l := 0
	r := len(ranges) - 1

	for l <= r {
		m := (l + r) / 2

		if ranges[m][0] <= id && id <= ranges[m][1] {
			return true
		}
		if ranges[m][1] < id {
			l = m + 1
		}
		if ranges[m][0] > id {
			r = m - 1
		}
	}
	return false
}

func Part1(input string) int {
	ranges, ingredients := formatInput(input)

	res := 0
	for _, ingredient := range ingredients {
		if isFresh(ranges, ingredient) {
			res++
		}
	}
	return res
}

func Part2(input string) int {
	ranges, _ := formatInput(input)
	res := 0

	l := 0
	r := -1
	for _, ran := range ranges {
		if r < ran[0] {
			res += r - l + 1

			l = ran[0]
			r = ran[1]
		} else {
			r = max(r, ran[1])
		}
	}
	res += r - l + 1

	return res
}
