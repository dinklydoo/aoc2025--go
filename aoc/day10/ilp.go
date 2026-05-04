package day10

import (
	"aoc2025--go/aoc/utils"
	"math"
	"slices"
)

func nullity(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	var null []int

	pivot := 0
	for c := 0; c < m-1; c++ {
		if pivot == n || matrix[pivot][c] == 0 {
			if pivot > 0 {
				null = append(null, c)
			}
			continue
		}
		pivot++
	}
	return null
}

func check(matrix [][]int, null []int, state []int, solve []int) int {
	v := len(matrix[0]) - 1
	values := make([]int, v)
	found := make([]bool, v)

	for i, n := range null {
		values[n] = state[i]
		found[n] = true
	}
	for _, eq := range solve {
		rhs := matrix[eq][v] // what we need to sum to

		fix := -1
		for i := 0; i < v; i++ {
			if matrix[eq][i] != 0 {
				if found[i] {
					rhs -= values[i] * matrix[eq][i]
				} else {
					fix = i
				}
			}
		}

		if fix >= 0 {
			if rhs%matrix[eq][fix] != 0 {
				return math.MaxInt
			}

			values[fix] = rhs / matrix[eq][fix]
			if values[fix] < 0 {
				return math.MaxInt
			}

			found[fix] = true
		} else if rhs != 0 {
			return math.MaxInt
		}
	}

	res := 0
	for _, v := range values {
		res += v
	}
	return res
}

func topoSort(matrix [][]int, null []int) []int {
	v := len(matrix[0]) - 1
	solved := make([]bool, v)
	for _, n := range null {
		solved[n] = true
	}

	var queue []int
	for row := range matrix {
		queue = append(queue, row)
	}

	var sorted []int

	for len(queue) > 0 {
		row := queue[0]
		queue = queue[1:]

		var unk []int
		for i := 0; i < v; i++ {
			if matrix[row][i] != 0 && !solved[i] {
				unk = append(unk, i)
			}
			if len(unk) > 1 {
				queue = append(queue, row)
				break
			}
		}

		if len(unk) == 1 {
			solved[unk[0]] = true
			sorted = append(sorted, row)
		}
	}
	return sorted
}

func hash(v []int) uint64 {
	var h uint64 = 1469598103934665603
	const prime uint64 = 1099511628211

	for _, x := range v {
		h ^= uint64(x + 1)
		h *= prime
	}
	return h
}

func Solve(matrix [][]int) int {
	null := nullity(matrix)
	solveOrder := topoSort(matrix, null)

	res := math.MaxInt
	var queue [][]int
	queue = append(queue, make([]int, len(null)))

	exists := make(map[uint64]bool)

	for i := 0; i < utils.Pow(250, len(null)); i++ {
		top := queue[0]
		queue = queue[1:]
		exists[hash(top)] = true

		res = min(res, check(matrix, null, top, solveOrder))

		for i := range null {
			top[i]++
			h := hash(top)
			if !exists[h] {
				temp := slices.Clone(top)
				queue = append(queue, temp)
				exists[h] = true
			}
			top[i]--
		}
	}
	return res
}
