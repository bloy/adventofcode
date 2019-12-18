package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
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

// FixedGrid is a fixed-size grid
type FixedGrid struct {
	values []rune
	Size   Point
}

// NewFixedGrid creates a fixed grid from a newline-separated string
func NewFixedGrid(str string) *FixedGrid {
	lines := strings.Split(str, "\n")
	size := Point{X: len(lines[0]), Y: len(lines)}
	values := make([]rune, 0)
	for _, line := range lines {
		values = append(values, []rune(line)...)
	}
	return &FixedGrid{
		values: values,
		Size:   size,
	}
}

func (g *FixedGrid) pointIndex(p Point) int {
	return p.Y*g.Size.X + p.X
}

// GetPoint gets the value at point p
func (g *FixedGrid) GetPoint(p Point) rune {
	return g.values[g.pointIndex(p)]
}

// SetPoint sets the point specified and returns the previous value
func (g *FixedGrid) SetPoint(p Point, value rune) rune {
	num := g.pointIndex(p)
	old := g.values[num]
	g.values[num] = value
	return old
}

// String implements the Stringer interface
func (g *FixedGrid) String() string {
	b := strings.Builder{}
	for y := 0; y < g.Size.Y; y++ {
		fmt.Fprint(&b, string(g.values[y*g.Size.X:(y+1)*g.Size.X]), "\n")
	}
	s := b.String()
	return s[:len(s)-1]
}

// Grid is a sparse grid of runes
type Grid struct {
	values             map[Point]rune
	blank              rune
	minPoint, maxPoint Point
	runeColor          map[rune]*color.Color
}

// NewGrid creates a new grid
func NewGrid() *Grid {
	g := &Grid{
		values:    make(map[Point]rune),
		runeColor: make(map[rune]*color.Color),
		blank:     '.',
	}
	return g
}

// NewGridFromInput takes a multiline string, building a grid out of it
func NewGridFromInput(in string) *Grid {
	g := NewGrid()
	for y, line := range strings.Split(in, "\n") {
		for x, r := range line {
			p := Point{X: x, Y: y}
			g.SetPoint(p, r)
		}
	}
	return g
}

// Copy copies a grid into new memory
func (g *Grid) Copy() *Grid {
	newg := NewGrid()
	newg.blank = g.blank
	newg.minPoint = g.minPoint
	newg.maxPoint = g.maxPoint
	for k, v := range g.values {
		newg.values[k] = v
	}
	for k, v := range g.runeColor {
		newg.runeColor[k] = v
	}
	return newg
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

// AddRuneColor adds a color that a rune should be
func (g *Grid) AddRuneColor(r rune, c *color.Color) {
	g.runeColor[r] = c
}

// ColorPrint prints in color directly to stdout
func (g *Grid) ColorPrint() {
	min, max := g.Bounds()
	defColor := color.New(color.Reset)
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			c := g.GetPoint(Point{X: x, Y: y})
			attrs, ok := g.runeColor[c]
			if !ok {
				attrs = defColor
			}
			attrs.Print(string(c))
		}
		defColor.Print("\n")
	}
}
