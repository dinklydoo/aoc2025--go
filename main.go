package main

import (
	"aoc2025--go/aoc/day01"
	"aoc2025--go/aoc/day02"
	"aoc2025--go/aoc/day03"
	"aoc2025--go/aoc/day04"
	"aoc2025--go/aoc/day05"
	"aoc2025--go/aoc/day06"
	"aoc2025--go/aoc/day07"
	"aoc2025--go/aoc/day08"
	"aoc2025--go/aoc/day09"
	"aoc2025--go/aoc/day10"
	"aoc2025--go/aoc/day11"
	"aoc2025--go/aoc/day12"
	"flag"
	"fmt"
	"os"
)

type dayFunc struct {
	part1 func(string) int
	part2 func(string) int
}

func main() {
	day := flag.Int("day", 1, "which day to run")
	part := flag.Int("part", 1, "which part to run")
	flag.Parse()

	input, err := os.ReadFile(fmt.Sprintf("input/day%02d.txt", *day))
	if err != nil {
		panic(err)
	}

	var days = map[int]dayFunc{
		1:  {day01.Part1, day01.Part2},
		2:  {day02.Part1, day02.Part2},
		3:  {day03.Part1, day03.Part2},
		4:  {day04.Part1, day04.Part2},
		5:  {day05.Part1, day05.Part2},
		6:  {day06.Part1, day06.Part2},
		7:  {day07.Part1, day07.Part2},
		8:  {day08.Part1, day08.Part2},
		9:  {day09.Part1, day09.Part2},
		10: {day10.Part1, day10.Part2},
		11: {day11.Part1, day11.Part2},
		12: {day12.Part1, day12.Part2},
	}

	run(days[*day].part1, days[*day].part2, *part, string(input))
}

func run(p1 func(string) int, p2 func(string) int, part int, input string) {
	if part == 1 {
		fmt.Println(p1(input))
	} else {
		fmt.Println(p2(input))
	}
}
