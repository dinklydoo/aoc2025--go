package day10

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type factory struct {
	config  []bool
	buttons [][]int
	joltage []int
}

func formatInput(input string) []factory {
	lines := strings.Split(input, "\n")
	factories := make([]factory, len(lines))

	for i, line := range lines {
		tokens := strings.Split(line, " ")
		factory := &factories[i]
		for _, token := range tokens {
			delim := token[1 : len(token)-1]

			switch token[0] {
			case '[':
				for _, machine := range delim {
					if machine == '#' {
						factory.config = append(factory.config, true)
					} else {
						factory.config = append(factory.config, false)
					}
				}
			case '(':
				args := strings.Split(delim, ",")
				var temp []int
				for _, arg := range args {
					val, _ := strconv.Atoi(arg)
					temp = append(temp, val)
				}
				factory.buttons = append(factory.buttons, temp)
			case '{':
				jolts := strings.Split(delim, ",")
				for _, jolt := range jolts {
					val, _ := strconv.Atoi(jolt)
					factory.joltage = append(factory.joltage, val)
				}
			}
		}
	}
	return factories
}

func Part1(input string) int {
	factories := formatInput(input)
	// this is probs NP-hard (SAT-like) so brute force with power set, our input seems limited by like 10 button configurations so pre chill
	res := 0
	for i := range factories {
		factory := &factories[i]

		buttons := len(factory.buttons)
		machines := len(factory.config)

		temp := math.MaxInt

		bitmask := (1 << buttons) - 1
		for pset := 0; pset <= bitmask; pset++ {
			toggle := make([]bool, machines)

			count := 0
			sel := 1
			for i := 0; i < buttons; i++ {
				if pset&sel > 0 {
					count++
					for _, sw := range factory.buttons[i] {
						toggle[sw] = toggle[sw] != true
					}
				}
				sel <<= 1
			}

			can := true
			for i := 0; i < machines; i++ {
				if toggle[i] != factory.config[i] {
					can = false
					break
				}
			}
			if can {
				temp = min(temp, count)
			}
		}
		res += temp
	}
	return res
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

func bfs(factory *factory) int {
	visited := make(map[uint64]bool)
	var queue [][]int
	press := 0
	count := 1
	queue = append(queue, slices.Clone(factory.joltage))
	for len(queue) > 0 {
		if count == 0 {
			press++
			count = len(queue)
		}
		top := queue[0]
		queue = queue[1:]

		good := true
		for _, jolt := range top {
			if jolt > 0 {
				good = false
				break
			}
		}
		if good {
			return press
		}

		for _, button := range factory.buttons {
			good := true
			for _, b := range button {
				if top[b] == 0 {
					good = false
					break
				}
			}
			if good {
				temp := slices.Clone(top)
				for _, b := range button {
					temp[b]--
				}
				hash := hash(temp)
				if visited[hash] {
					continue
				} else {
					visited[hash] = true
				}
				queue = append(queue, temp)
			}
		}
		count--
	}
	return -1
}

// TODO too slow, need Gaussian Elim

func Part2(input string) int {
	factories := formatInput(input)
	res := 0
	for _, factory := range factories {
		res += bfs(&factory)
		fmt.Println(res)
	}
	return res
}
