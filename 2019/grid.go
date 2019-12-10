package main

import (
	"fmt"
	"strings"
)

// 2d Direction vectors
var (
	North = Point{X: 0, Y: -1}
	South = Point{X: 0, Y: 1}
	East  = Point{X: 1, Y: 0}
	West  = Point{X: -1, Y: 0}
	Up    = North
	Down  = South
	Left  = West
	Right = East
)

// Point represents a point (or vector) in 2d space
type Point struct {
	X, Y int
}

// Distance returns the manhattan distance between two points
func (p Point) Distance(p2 Point) int {
	xdist := p.X - p2.X
	ydist := p.Y - p2.Y
	if xdist < 0 {
		xdist *= -1
	}
	if ydist < 0 {
		ydist *= -1
	}
	return xdist + ydist
}

// Add adds the vector v to the point, returning a new point
func (p Point) Add(v Point) Point {
	return Point{
		X: p.X + v.X,
		Y: p.Y + v.Y,
	}
}

// Grid is a sparse grid of runes
type Grid struct {
	values             map[Point]rune
	blank              rune
	minPoint, maxPoint Point
}

// NewGrid creates a new grid
func NewGrid() *Grid {
	g := &Grid{
		values: make(map[Point]rune),
		blank:  '.',
	}
	return g
}

// SetBlank sets the blank rune for a grid
func (g *Grid) SetBlank(r rune) {
	g.blank = r
}

// SetPoint sets the point specified and returns the previous value if any
func (g *Grid) SetPoint(p Point, value rune) rune {
	if p.X < g.minPoint.X {
		g.minPoint.X = p.X
	}
	if p.Y < g.minPoint.Y {
		g.minPoint.Y = p.Y
	}
	if p.X > g.maxPoint.X {
		g.maxPoint.X = p.X
	}
	if p.Y > g.maxPoint.Y {
		g.maxPoint.Y = p.Y
	}
	old := g.values[p]
	g.values[p] = value
	return old
}

// GetPoint gets the value of the grid at p
func (g *Grid) GetPoint(p Point) rune {
	r, ok := g.values[p]
	if ok {
		return r
	}
	return g.blank
}

// Bounds returns the minimum and maximum points that bound this grid
func (g *Grid) Bounds() (minPoint, maxPoint Point) {
	return g.minPoint, g.maxPoint
}

// String implements the Stringer interface
func (g *Grid) String() string {
	b := strings.Builder{}
	min, max := g.Bounds()
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			p := Point{X: x, Y: y}
			fmt.Fprint(&b, string(g.GetPoint(p)))
		}
		fmt.Fprint(&b, "\n")
	}
	return b.String()
}
