package day10

import (
	"aoc2025--go/aoc/utils"
	"math/big"
)

func normalize(pivotRow []big.Rat, factor big.Rat) {
	var inv big.Rat
	inv.Inv(&factor)
	for i := 0; i < len(pivotRow); i++ {
		pivotRow[i].Mul(&pivotRow[i], &inv)
	}
}

func eliminate(pivotRow []big.Rat, row []big.Rat, factor big.Rat) {
	var tmp big.Rat
	for i := 0; i < len(row); i++ {
		tmp.Mul(&pivotRow[i], &factor)
		row[i].Sub(&row[i], &tmp)
	}
}

func normalizeMatrix(matrix [][]big.Rat) [][]int {
	n := len(matrix)
	m := len(matrix[0])

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for r := 0; r < n; r++ {
		lcm := 1
		for c := 0; c < m; c++ {
			temp := matrix[r][c].Denom().Int64()
			lcm = utils.LCM(lcm, int(temp))
		}
		for c := 0; c < m; c++ {
			num := matrix[r][c].Num().Int64()
			den := matrix[r][c].Denom().Int64()
			res[r][c] = int(num * int64(lcm) / den)
		}
	}
	return res
}

func Reduce(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])

	matrix := make([][]big.Rat, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]big.Rat, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j].SetInt64(int64(a[i][j]))
		}
	}

	pivot := 0 // pivot row
	for c := 0; c < m-1; c++ {
		if pivot == n {
			break
		}

		if matrix[pivot][c].Sign() == 0 {
			found := false
			for r := pivot + 1; r < n; r++ {
				if matrix[r][c].Sign() != 0 {
					found = true
					matrix[r], matrix[pivot] = matrix[pivot], matrix[r]
					break
				}
			}
			if !found {
				continue
			}
		}
		var one big.Rat
		one.SetInt64(1)
		if matrix[pivot][c].Cmp(&one) != 0 {
			utils.Assert(matrix[pivot][c].Sign() != 0, "row pivot should be non-zero")

			var factor big.Rat
			factor.Set(&matrix[pivot][c])
			normalize(matrix[pivot], factor)
		}
		utils.Assert(matrix[pivot][c].Cmp(&one) == 0, "row pivot should be one")
		for r := pivot + 1; r < n; r++ {
			// eliminate pivot values below
			if matrix[r][c].Sign() != 0 {
				var factor big.Rat
				factor.Set(&matrix[r][c])
				eliminate(matrix[pivot], matrix[r], factor)
			}
		}
		pivot++
	}
	// normalize ratiojnal to int
	imatrix := normalizeMatrix(matrix)
	IRF(imatrix)
	return imatrix
}

func simplify(pivotRow []int, row []int, factor int) {
	for i := 0; i < len(pivotRow); i++ {
		row[i] -= pivotRow[i] * factor
	}
}

// IRF : reduce the matrix to integer row form, more relaxed variant of RREF
func IRF(a [][]int) {
	n := len(a)
	m := len(a[0])

	pivot := -1
	for c := 0; c < m-1; c++ {
		found := false
		for r := pivot + 1; r < n; r++ {
			if a[r][c] == 1 {
				found = true
				pivot = r
				break
			}
		}
		if !found {
			continue
		}
		for r := pivot - 1; r >= 0; r-- {
			if a[r][c] != 0 {
				simplify(a[pivot], a[r], a[r][c])
			}
		}
	}
}
