package main

import (
	"fmt"
	"strings"
)

type Tile int8

const (
	OPEN  Tile = iota // square is an open acre
	TREES             // square is a tree filled acre
	YARD              // square is a lumberyard
)

func (s Tile) String() string {
	switch s {
	case OPEN:
		return "."
	case TREES:
		return "|"
	case YARD:
		return "#"
	default:
		return " "
	}
}

type Area struct {
	Tiles []Tile
	Xsize int
	Ysize int
}

func (a *Area) String() string {
	var out strings.Builder
	for y := 0; y < a.Ysize; y++ {
		for x := 0; x < a.Xsize; x++ {
			fmt.Fprint(&out, a.Tiles[y*a.Xsize+x])
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
