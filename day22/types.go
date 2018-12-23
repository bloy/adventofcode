package main

import (
	"fmt"
	"strings"
)

type Point struct {
	X, Y int
}

var NORTH Point = Point{0, -1}
var EAST Point = Point{1, 0}
var SOUTH Point = Point{0, 1}
var WEST Point = Point{-1, 0}

var DIRECTIONS [4]Point = [4]Point{NORTH, EAST, SOUTH, WEST}

type RegionType int8

const (
	UNKNOWN RegionType = iota
	ROCKY
	WET
	NARROW
)

type Item int8

const (
	NEITHER Item = iota
	TORCH
	CLIMB
)

func (i Item) String() string {
	switch i {
	case NEITHER:
		return "Neither"
	case TORCH:
		return "Torch"
	case CLIMB:
		return "Climbing Gear"
	default:
		return "Unknown"
	}
}

var EQUIPMENT [3]Item = [3]Item{NEITHER, TORCH, CLIMB}

var ItemNotAllowed map[RegionType]Item = map[RegionType]Item{
	ROCKY:  NEITHER,
	WET:    TORCH,
	NARROW: CLIMB,
}

type Region struct {
	Point
	Type         RegionType
	GeoIndex     int
	ErosionLevel int
}

func (r *Region) Risk() int {
	switch r.Type {
	case ROCKY:
		return 0
	case WET:
		return 1
	case NARROW:
		return 2
	}
	return 0
}

type Cave struct {
	Depth   int
	Regions map[Point]*Region
	Target  Point
}

func initRegion(r *Region, cave *Cave) {
	if (r.X == 0 && r.Y == 0) || (r.X == cave.Target.X && r.Y == cave.Target.Y) {
		r.GeoIndex = 0
	} else if r.Y == 0 {
		r.GeoIndex = r.X * 16807
	} else if r.X == 0 {
		r.GeoIndex = r.Y * 48271
	} else {
		r1 := cave.Region(r.X-1, r.Y)
		r2 := cave.Region(r.X, r.Y-1)
		r.GeoIndex = r1.ErosionLevel * r2.ErosionLevel
	}
	r.ErosionLevel = (r.GeoIndex + cave.Depth) % 20183
	switch r.ErosionLevel % 3 {
	case 0:
		r.Type = ROCKY
	case 1:
		r.Type = WET
	case 2:
		r.Type = NARROW
	}
}

func (c *Cave) Region(x, y int) *Region {
	r := c.Regions[Point{x, y}]
	if r == nil {
		r = &Region{}
		r.X = x
		r.Y = y
		c.Regions[Point{x, y}] = r
		initRegion(r, c)
	}
	return r
}

func NewCave(depth, targetx, targety int) *Cave {
	c := Cave{}
	c.Depth = depth
	c.Regions = make(map[Point]*Region)
	c.Target = Point{targetx, targety}
	return &c
}

func (c *Cave) String() string {
	var str strings.Builder
	for y := 0; y <= c.Target.Y; y++ {
		for x := 0; x <= c.Target.X; x++ {
			r := c.Region(x, y)
			if r.X == 0 && r.Y == 0 {
				str.WriteString("M")
			} else if r.X == c.Target.X && r.Y == c.Target.Y {
				str.WriteString("T")
			} else if r.Type == ROCKY {
				str.WriteString(".")
			} else if r.Type == NARROW {
				str.WriteString("|")
			} else if r.Type == WET {
				str.WriteString("=")
			}
		}
		str.WriteString("\n")
	}
	return str.String()
}

type Node struct {
	Point
	Minutes  int
	Equipped Item
}

func (n *Node) String() string {
	return fmt.Sprintf("(<%d, %d> %d min, %v)", n.X, n.Y, n.Minutes, n.Equipped)
}

type SeenNode struct {
	Point
	Equipped Item
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Minutes < pq[j].Minutes
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
