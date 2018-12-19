package main

import (
	"fmt"
	"strings"
)

type Square int8

const (
	SAND    Square = iota // square is filled with dry sand
	CLAY                  // square is filled with clay
	FLOWING               // square is filled with flowing water and sand
	SETTLED               // square is filled with settled water and sand
	SOURCE                // water source
)

func (s Square) String() string {
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

func (p Point) up() Point {
	return Point{p.x, p.y - 1}
}

func (p Point) down() Point {
	return Point{p.x, p.y + 1}
}

func (p Point) left() Point {
	return Point{p.x - 1, p.y}
}

func (p Point) right() Point {
	return Point{p.x + 1, p.y}
}

type Area struct {
	Squares            map[Point]Square
	MinPoint, MaxPoint Point
}

func (a *Area) String() string {
	var out strings.Builder
	for y := a.MinPoint.y; y <= a.MaxPoint.y; y++ {
		for x := a.MinPoint.x; x <= a.MaxPoint.x; x++ {
			fmt.Fprint(&out, a.Squares[Point{x, y}])
		}
		out.WriteString("\n")
	}
	return out.String()
}

func NewArea() *Area {
	a := &Area{}
	a.Squares = make(map[Point]Square)
	return a
}
