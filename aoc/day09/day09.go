package day09

import (
	"aoc2025--go/aoc/utils"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p Point) Rectangle(o Point) int {
	return (1 + utils.Abs(o.x-p.x)) * (1 + utils.Abs(o.y-p.y))
}

func formatInput(input string) []Point {
	var points []Point
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		point := strings.Split(line, ",")
		x, _ := strconv.Atoi(point[0])
		y, _ := strconv.Atoi(point[1])

		points = append(points, Point{x, y})
	}
	return points
}

func Part1(input string) int {
	points := formatInput(input)

	res := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pi := &points[i]
			pj := &points[j]
			res = max(res, pi.Rectangle(*pj))
		}
	}
	return res
}

const VERTICAL = 0
const HORIZONTAL = 1

type Edge struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func order(a, b *Point) (*Point, *Point) {
	if a.x == b.x { // vertical
		if a.y < b.y {
			return a, b
		}
		return b, a
	}
	if a.x < b.x { // horizontal
		return a, b
	}
	return b, a
}

func mergeEdges(edges [2][]Edge) [2][]Edge {
	slices.SortFunc(edges[VERTICAL], func(a, b Edge) int {
		if a.x1 == b.x1 {
			return a.y1 - b.y1
		}
		return a.x1 - b.x1
	})
	slices.SortFunc(edges[HORIZONTAL], func(a, b Edge) int {
		if a.y1 == b.y1 {
			return a.x1 - b.x1
		}
		return a.y1 - b.y1
	})

	var merged [2][]Edge

	merge := func(list []Edge, vertical bool) []Edge {
		var res []Edge

		var start, end, fixed int

		if vertical {
			fixed = list[0].x1
			start = list[0].y1
			end = list[0].y2
		} else {
			fixed = list[0].y1
			start = list[0].x1
			end = list[0].x2
		}

		for i := 1; i < len(list); i++ {
			e := list[i]

			var s, e2, f int
			var sameGroup bool

			if vertical {
				sameGroup = e.x1 == fixed
				s, e2, f = e.y1, e.y2, e.x1
			} else {
				sameGroup = e.y1 == fixed
				s, e2, f = e.x1, e.x2, e.y1
			}

			if sameGroup && s <= end {
				end = max(end, e2)
				continue
			}

			if vertical {
				res = append(res, Edge{fixed, fixed, start, end})
			} else {
				res = append(res, Edge{start, end, fixed, fixed})
			}

			fixed = f
			start = s
			end = e2
		}
		if vertical {
			res = append(res, Edge{fixed, fixed, start, end})
		} else {
			res = append(res, Edge{start, end, fixed, fixed})
		}
		return res
	}

	merged[VERTICAL] = merge(edges[VERTICAL], true)
	merged[HORIZONTAL] = merge(edges[HORIZONTAL], false)

	return merged
}

func Part2(input string) int {
	points := formatInput(input)

	var edges [2][]Edge
	n := len(points)
	for i := 0; i < n; i++ {
		p1, p2 := order(&points[i], &points[(i-1+n)%n])
		if p1.x == p2.x {
			edges[VERTICAL] = append(edges[VERTICAL], Edge{p1.x, p2.x, p1.y, p2.y})
		} else {
			edges[HORIZONTAL] = append(edges[HORIZONTAL], Edge{p1.x, p2.x, p1.y, p2.y})
		}
	}
	edges = mergeEdges(edges)

	res := 0
	for i := 0; i < len(points); i++ {
	outerLoop:
		for j := i + 1; j < len(points); j++ {
			x1 := min(points[i].x, points[j].x)
			x2 := max(points[i].x, points[j].x)
			y1 := min(points[i].y, points[j].y)
			y2 := max(points[i].y, points[j].y)

			size := points[i].Rectangle(points[j])
			if size < res {
				continue
			}

			start, _ := slices.BinarySearchFunc(edges[VERTICAL], x1, func(edge Edge, x int) int {
				return edge.x1 - x
			})

			for start < len(edges[VERTICAL]) && edges[VERTICAL][start].x1 <= x2 {
				edge := edges[VERTICAL][start]

				if edge.y2 > y1 && edge.y1 < y2 { // overlap exists

					if edge.x1 > x1 && edge.x2 < x2 { // slice-through
						continue outerLoop
					}
				}
				start++
			}

			start, _ = slices.BinarySearchFunc(edges[HORIZONTAL], y1, func(edge Edge, y int) int {
				return edge.y1 - y
			})

			for start < len(edges[HORIZONTAL]) && edges[HORIZONTAL][start].y1 <= y2 {
				edge := edges[HORIZONTAL][start]

				if edge.x2 > x1 && edge.x1 < x2 { // overlap exists

					if edge.y1 > y1 && edge.y2 < y2 { // slice-through
						continue outerLoop
					}
				}
				start++
			}
			res = max(res, size)
		}
	}
	return res
}
