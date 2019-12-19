package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

func init() {
	AddSolution(18, solveDay18)
}

type day18OuterState struct {
	place rune
	keys  int
}

type day18State struct {
	pos      Point           // pos is for the "inner" A*
	state    day18OuterState // state is for the "outer" A*
	priority float64
	index    int
}

type day18Queue []*day18State

func (q day18Queue) Len() int { return len(q) }

func (q day18Queue) Less(i, j int) bool {
	return q[i].priority > q[j].priority
}

func (q day18Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *day18Queue) Push(x interface{}) {
	n := len(*q)
	item := x.(*day18State)
	item.index = n
	*q = append(*q, item)
}

func (q *day18Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1
	old[n-1] = nil
	*q = old[:n-1]
	return item
}

type day18Grid struct {
	FixedGrid
	keyBits   map[rune]int
	doorBits  map[rune]int
	locations map[rune]Point
	allKeys   int
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
	mapgrid = `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`
	steps, err := day18Part1(mapgrid)
	pr.CheckError(err)
	pr.ReportPart(steps)
}

func day18ShortestPath(grid *day18Grid, keys int, from, to Point) (shortestPath []Point, found bool) {
	seen := make(map[Point]bool)
	openHeap := make(day18Queue, 0)
	heap.Init(&openHeap)
	cameFrom := make(map[Point]Point)
	gScore := make(map[Point]float64)
	fScore := make(map[Point]float64)
	h := func(p Point) float64 {
		return math.Hypot(float64(to.X-p.X), float64(to.Y-p.Y))
	}
	gScore[from] = 0
	fScore[from] = gScore[from] + h(from)
	heap.Push(&openHeap, &day18State{pos: from, priority: fScore[from]})
	seen[from] = true
	for len(openHeap) > 0 {
		pos := heap.Pop(&openHeap).(*day18State).pos
		if pos == to {
			path := make([]Point, 0)
			p := pos
			for p != from {
				path = append([]Point{p}, path...)
				p = cameFrom[p]
			}
			return path, true
		}
		for _, p := range grid.AdjacentPoints(pos) {
			if seen[p] {
				continue
			}
			seen[p] = true
			cameFrom[p] = pos
			gScore[p] = gScore[pos] + 1
			fScore[p] = gScore[p] + h(p)
			c := grid.GetPoint(p)
			if c == '#' {
				continue
			}
			if c >= 'A' && c <= 'Z' && (keys&grid.doorBits[c] == 0) {
				// run into a locked door
				continue
			}
			if c >= 'a' && c <= 'z' && (keys&grid.keyBits[c] == 0) && grid.locations[c] != to {
				// need to find a path that picks up this key first
				continue
			}
			heap.Push(&openHeap, &day18State{pos: p, priority: fScore[p]})
		}
	}
	return nil, false
}

// Box holds the same information as 2 points
type Box struct {
	X1, Y1, X2, Y2 int
}

func day18Part1(mapgrid string) (steps int, err error) {
	grid := &day18Grid{}
	grid.FixedGrid = *NewFixedGrid(mapgrid)
	grid.keyBits = make(map[rune]int)
	grid.doorBits = make(map[rune]int)
	grid.locations = make(map[rune]Point)
	for y := 0; y < grid.Size.Y; y++ {
		for x := 0; x < grid.Size.X; x++ {
			p := Point{X: x, Y: y}
			c := grid.GetPoint(p)
			if c == '@' {
				grid.locations[c] = p
			} else if c >= 'a' && c <= 'z' {
				grid.keyBits[c] = 1 << (c - 'a')
				grid.locations[c] = p
				grid.allKeys |= grid.keyBits[c]
			} else if c >= 'A' && c <= 'Z' {
				grid.doorBits[c] = 1 << (c - 'A')
				grid.locations[c] = p
			}
		}
	}
	fmt.Println(grid)
	keys := make([]rune, 0)
	for k := range grid.keyBits {
		keys = append(keys, k)
	}

	seen := make(map[day18OuterState]bool)
	openHeap := make(day18Queue, 0)
	heap.Init(&openHeap)
	cameFrom := make(map[rune]rune)
	pathFrom := make(map[Box][]Point)
	gScore := make(map[day18OuterState]float64)
	fScore := make(map[day18OuterState]float64)
	h := func(s day18OuterState) float64 {
		count := 0
		for _, k := range keys {
			if s.keys|grid.keyBits[k] == 0 {
				count++
			}
		}
		return float64(count)
	}
	start := day18OuterState{place: '@', keys: 0}
	gScore[start] = 0
	fScore[start] = gScore[start] + h(start)
	seen[start] = true
	heap.Push(&openHeap, &day18State{state: start, priority: fScore[start]})
	for len(openHeap) > 0 {
		state := heap.Pop(&openHeap).(*day18State).state
		if state.keys == grid.allKeys {
			loc := state.place
			path := ""
			for loc != '@' {
				path = string(loc) + path
				loc = cameFrom[loc]
			}
			fmt.Println(path)
			return int(gScore[state]), nil
		}
		for _, k := range keys {
			if state.place == k {
				continue
			}
			if state.keys&grid.keyBits[k] == grid.keyBits[k] {
				continue // already have this key
			}
			cur := grid.locations[state.place]
			dest := grid.locations[k]
			pathlen := 0
			newState := day18OuterState{place: k, keys: state.keys | grid.keyBits[k]}
			if seen[newState] {
				continue
			}
			seen[newState] = true
			if path, found := day18ShortestPath(grid, state.keys, cur, dest); found {
				pathFrom[Box{X1: cur.X, Y1: cur.Y, X2: dest.X, Y2: dest.Y}] = path
				pathlen = len(path)
			} else {
				continue // unable to find a path to that key yet
			}
			cameFrom[newState.place] = state.place
			gScore[newState] = gScore[state] + float64(pathlen)
			fScore[newState] = gScore[newState] + h(newState)
			heap.Push(&openHeap, &day18State{state: newState, priority: fScore[newState]})
		}
	}
	return 0, fmt.Errorf("Unable to find a path")
}
