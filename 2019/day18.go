package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

func init() {
	AddSolution(18, solveDay18)
}

type day18State struct {
	pos  Point
	keys int
}

func (s day18State) NewPos(p Point) day18State {
	return day18State{pos: p, keys: s.keys}
}

func (s day18State) NewKeys(keys int) day18State {
	return day18State{pos: s.pos, keys: keys}
}

type day18Path []day18State

func (p day18Path) Copy() day18Path {
	n := make(day18Path, len(p))
	for i := range p {
		n[i] = p[i]
	}
	return n
}

func solveDay18(pr *PuzzleRun) {
	lines := []string{}
	scanner := bufio.NewScanner(pr.InFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		pr.CheckError(err)
	}
	pr.ReportLoad()
	mapgrid := strings.Join(lines, "\n")
	mapgrid = `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`
	steps, err := day18Part1(mapgrid)
	pr.CheckError(err)
	pr.ReportPart(steps)
}

func day18Part1(mapgrid string) (steps int, err error) {
	grid := NewFixedGrid(mapgrid)
	var pos Point
	keyBits := make(map[rune]int)
	doorBits := make(map[rune]int)
	var allKeys int
	for y := 0; y < grid.Size.Y; y++ {
		for x := 0; x < grid.Size.X; x++ {
			p := Point{X: x, Y: y}
			c := grid.GetPoint(p)
			if c == '@' {
				pos = p
			} else if c >= 'a' && c <= 'z' {
				keyBits[c] = 1 << (c - 'a')
				allKeys |= keyBits[c]
			} else if c >= 'A' && c <= 'Z' {
				doorBits[c] = 1 << (c - 'A')
			}
		}
	}
	fmt.Println(grid)

	seen := make(map[day18State]bool)
	current := day18State{pos: pos, keys: 0}
	path := day18Path{current}
	stack := make([]day18Path, 1)
	seen[current] = true
	stack[0] = path.Copy()
	var shortestLen int = math.MaxInt32
	for len(stack) > 0 {
		path = stack[0]
		stack = stack[1:] // pop the front off the stack
		if len(path) > shortestLen {
			continue
		}
		current = path[len(path)-1]
		seen[current] = true
		location := grid.GetPoint(current.pos)
		if location >= 'a' && location <= 'z' {
			current = current.NewKeys(current.keys | keyBits[location])
			seen[current] = true
		}
		if allKeys == current.keys {
			if len(path) < shortestLen {
				shortestLen = len(path)
			}
			continue // no need to continue further on this path
		}
		for _, nextPoint := range grid.AdjacentPoints(current.pos) {
			location = grid.GetPoint(nextPoint)
			if location == '#' {
				continue
			}
			if location >= 'A' && location <= 'Z' && (current.keys&doorBits[location] == 0) {
				continue // encountered an impassible door
			}
			next := current.NewPos(nextPoint)
			if seen[next] {
				continue
			}
			nextPath := append(path.Copy(), next)
			stack = append(stack, nextPath)
		}
	}
	return shortestLen - 1, nil
}
