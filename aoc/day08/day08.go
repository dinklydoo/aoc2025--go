package day08

import (
	"container/heap"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

func (p Point) distance(o Point) int {
	dx := o.x - p.x
	dy := o.y - p.y
	dz := o.z - p.z
	return dx*dx + dy*dy + dz*dz
}

func formatInput(input string) []Point {
	lines := strings.Split(input, "\n")
	points := make([]Point, len(lines))

	for i, line := range lines {
		point := strings.Split(line, ",")

		X, _ := strconv.Atoi(point[0])
		Y, _ := strconv.Atoi(point[1])
		Z, _ := strconv.Atoi(point[2])

		points[i] = Point{X, Y, Z}
	}
	return points
}

type Connection struct {
	from     int
	to       int
	distance int
}

func Part1(input string) int {
	points := formatInput(input)

	var connections []Connection

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d := points[i].distance(points[j])
			connections = append(connections, Connection{i, j, d})
		}
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})

	ds := initDSU(len(points))
	for i := 0; i < 1000; i++ {
		conn := &connections[i]
		ds.union(conn.from, conn.to)
	}

	var setSize = map[int]int{}
	for _, set := range ds.dsu {
		parent := ds.find(set)
		setSize[parent] = ds.size[parent]
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, size := range setSize {
		heap.Push(&pq, size)
	}

	res := 1
	for i := 0; i < 3; i++ {
		top := heap.Pop(&pq)
		res *= top.(int)
	}
	return res
}

func Part2(input string) int {
	points := formatInput(input)

	var connections []Connection

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d := points[i].distance(points[j])
			connections = append(connections, Connection{i, j, d})
		}
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})

	var res int

	ds := initDSU(len(points))
	i := 0
	for ds.sets > 1 {
		conn := &connections[i]
		ds.union(conn.from, conn.to)

		if ds.sets == 1 { // this connection joins all
			fx := points[conn.from].x
			tx := points[conn.to].x

			res = fx * tx
			break
		}
		i++
	}
	return res
}
