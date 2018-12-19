package main

import (
	"fmt"
	"strings"
)

type Tile int8

const (
	SAND    Tile = iota // square is filled with dry sand
	CLAY                // square is filled with clay
	FLOWING             // square is filled with flowing water and sand
	SETTLED             // square is filled with settled water and sand
	SOURCE              // water source
)

func (s Tile) String() string {
	switch s {
	case SAND:
		return "."
	case CLAY:
		return "#"
	case FLOWING:
		return "|"
	case SETTLED:
		return "~"
	case SOURCE:
		return "+"
	default:
		return " "
	}
}

type Point struct {
	x, y int
}

func (p Point) Up() Point {
	return Point{p.x, p.y - 1}
}

func (p Point) Down() Point {
	return Point{p.x, p.y + 1}
}

func (p Point) Left() Point {
	return Point{p.x - 1, p.y}
}

func (p Point) Right() Point {
	return Point{p.x + 1, p.y}
}

type Area struct {
	Tiles              map[Point]Tile
	MinPoint, MaxPoint Point
}

func (a *Area) String() string {
	var out strings.Builder
	for y := a.MinPoint.y; y <= a.MaxPoint.y; y++ {
		for x := a.MinPoint.x; x <= a.MaxPoint.x; x++ {
			fmt.Fprint(&out, a.Tiles[Point{x, y}])
		}
		out.WriteString("\n")
	}
	return out.String()
}

func NewArea() *Area {
	a := &Area{}
	a.Tiles = make(map[Point]Tile)
	return a
}
