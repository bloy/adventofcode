package main

import (
	"log"
	"strconv"
	"strings"
)

type dataType []*point

type point struct {
	pos   [4]int
	links []*point
}

func (p point) distanceTo(other *point) int {
	total := 0
	for i := 0; i < 4; i++ {
		total += abs(p.pos[i] - other.pos[i])
	}
	return total
}

type seenMap map[[4]int]bool

func newPoint(x, y, z, t int) *point {
	return &point{[4]int{x, y, z, t}, make([]*point, 0)}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Number conversion error:", err)
	}
	return i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func readInput(inputText string) dataType {
	lines := strings.Split(inputText, "\n")
	data := make(dataType, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, ",")
		data[i] = newPoint(
			atoi(nums[0]),
			atoi(nums[1]),
			atoi(nums[2]),
			atoi(nums[3]),
		)
	}
	return data
}

func connectConstellations(data dataType) dataType {
	for i := range data {
		for j := range data {
			if i == j {
				continue
			}
			if data[i].distanceTo(data[j]) <= 3 {
				data[i].links = append(data[i].links, data[j])
			}
		}
	}
	return data
}

func recurseConnections(seen seenMap, p *point) seenMap {
	seen[p.pos] = true
	for i := range p.links {
		if seen[p.links[i].pos] {
			continue
		}
		seen = recurseConnections(seen, p.links[i])
	}
	return seen
}

func solve1(data dataType) int {
	data = connectConstellations(data)
	seen := make(seenMap)
	count := 0
	for i := range data {
		if !seen[data[i].pos] {
			count++
		}
		seen = recurseConnections(seen, data[i])
	}
	return count
}

func solve2(data dataType) int {
	return 0
}
