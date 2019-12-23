package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func init() {
	AddSolution(18, solveDay18)
}

const (
	day18test1 = `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`

	day18test2 = `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`

	day18test3 = `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`
)

func solveDay18(pr *PuzzleRun) {
	buf, err := ioutil.ReadAll(pr.InFile)
	pr.CheckError(err)
	mapgrid := strings.TrimSpace(string(buf))
	pr.logger.Println(mapgrid)
	pr.ReportLoad()
	steps, err := day18Part1(mapgrid)
	pr.CheckError(err)
	pr.ReportPart(steps)
	steps, err = day18Part2(mapgrid)
	pr.CheckError(err)
	pr.ReportPart(steps)
}

type day18Grid struct {
	FixedGrid
	keys    []rune
	places  map[rune]Point
	bits    map[rune]int
	allKeys int
}

type day18State struct {
	pos  Point
	keys int
}

func day18MakeGrid(mapgrid string) *day18Grid {
	grid := &day18Grid{}
	grid.FixedGrid = *NewFixedGrid(mapgrid)
	grid.places = make(map[rune]Point)
	grid.bits = make(map[rune]int)
	grid.keys = make([]rune, 0)
	for y := 0; y < grid.FixedGrid.Size.Y; y++ {
		for x := 0; x < grid.FixedGrid.Size.X; x++ {
			p := Point{X: x, Y: y}
			c := grid.GetPoint(p)
			if c == '@' || (c >= 'A' && c <= 'Z') || c >= 'a' && c <= 'z' {
				grid.places[c] = p
			}
			if c >= 'a' && c <= 'z' {
				grid.keys = append(grid.keys, c)
			}
		}
	}
	for i, k := range grid.keys {
		grid.bits[k] = 1 << i
		grid.bits[rune(strings.ToUpper(string(k))[0])] = 1 << i
		grid.allKeys = grid.allKeys | (1 << i)
	}
	return grid
}

func day18search(grid *day18Grid, start Point, haveKeys int) (steps int, err error) {
	dist := make(map[day18State]int)
	seen := make(map[day18State]bool)
	current := day18State{pos: start, keys: haveKeys}
	q := make([]day18State, 1)
	q[0] = current
	seen[current] = true
	dist[current] = 0
	found := false
	foundState := current
	for len(q) > 0 && !found {
		current := q[0]
		q = q[1:]
		for _, adj := range grid.AdjacentPoints(current.pos) {
			keys := current.keys
			c := grid.GetPoint(adj)
			if c == '#' {
				continue
			} else if c >= 'a' && c <= 'z' {
				keys = keys | grid.bits[c]
			} else if c >= 'A' && c <= 'Z' && (keys&grid.bits[c] == 0) {
				continue
			}
			next := day18State{pos: adj, keys: keys}
			if seen[next] {
				continue
			}
			seen[next] = true
			dist[next] = dist[current] + 1
			q = append(q, next)
			if next.keys == grid.allKeys {
				found = true
				foundState = next
				break
			}
		}
	}
	if !found {
		return 0, fmt.Errorf("No path found")
	}
	return dist[foundState], nil
}

func day18Part1(mapgrid string) (steps int, err error) {
	grid := day18MakeGrid(mapgrid)
	steps, err = day18search(grid, grid.places['@'], 0)
	return
}

func day18Part2(mapgrid string) (steps int, err error) {
	grid := day18MakeGrid(mapgrid)
	center := grid.places['@']
	grid.SetPoint(center, '#')
	grid.SetPoint(center.Add(North), '#')
	grid.SetPoint(center.Add(South), '#')
	grid.SetPoint(center.Add(East), '#')
	grid.SetPoint(center.Add(West), '#')

	steps = 0
	haveKeys := grid.allKeys
	for y := 0; y < center.Y; y++ {
		for x := 0; x < center.X; x++ {
			c := grid.GetPoint(Point{X: x, Y: y})
			if c >= 'a' && c <= 'z' {
				haveKeys ^= grid.bits[c]
			}
		}
	}
	quadSteps, err := day18search(grid, center.Add(Point{X: -1, Y: -1}), haveKeys)
	if err != nil {
		return 0, err
	}
	steps += quadSteps

	haveKeys = grid.allKeys
	for y := center.Y + 1; y < grid.Size.Y; y++ {
		for x := 0; x < center.X; x++ {
			c := grid.GetPoint(Point{X: x, Y: y})
			if c >= 'a' && c <= 'z' {
				haveKeys ^= grid.bits[c]
			}
		}
	}
	quadSteps, err = day18search(grid, center.Add(Point{X: -1, Y: 1}), haveKeys)
	if err != nil {
		return 0, err
	}
	steps += quadSteps

	haveKeys = grid.allKeys
	for y := 0; y < center.Y; y++ {
		for x := center.X + 1; x < grid.Size.X; x++ {
			c := grid.GetPoint(Point{X: x, Y: y})
			if c >= 'a' && c <= 'z' {
				haveKeys ^= grid.bits[c]
			}
		}
	}
	quadSteps, err = day18search(grid, center.Add(Point{X: 1, Y: -1}), haveKeys)
	if err != nil {
		return 0, err
	}
	steps += quadSteps

	haveKeys = grid.allKeys
	for y := center.Y + 1; y < grid.Size.Y; y++ {
		for x := center.X + 1; x < grid.Size.X; x++ {
			c := grid.GetPoint(Point{X: x, Y: y})
			if c >= 'a' && c <= 'z' {
				haveKeys ^= grid.bits[c]
			}
		}
	}
	quadSteps, err = day18search(grid, center.Add(Point{X: 1, Y: 1}), haveKeys)
	if err != nil {
		return 0, err
	}
	steps += quadSteps

	return
}
