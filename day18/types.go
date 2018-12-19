package main

import (
	"fmt"
	"strings"
)

type Tile int8

type TileList []Tile

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
	Tiles TileList
	Xsize int
	Ysize int
}

func (a *Area) String() string {
	var out strings.Builder
	for y := 0; y < a.Ysize; y++ {
		for x := 0; x < a.Xsize; x++ {
			fmt.Fprint(&out, a.TileAt(x, y))
		}
		out.WriteString("\n")
	}
	return out.String()
}

func (a *Area) TileAt(x, y int) Tile {
	return a.Tiles[y*a.Xsize+x]
}

func (a *Area) SetTileAt(x, y int, t Tile) Tile {
	var prev Tile = a.TileAt(x, y)
	a.Tiles[y*a.Xsize+x] = t
	return prev
}

func (a *Area) Value() int {
	var yardCount, treeCount int
	for y := 0; y < a.Ysize; y++ {
		for x := 0; x < a.Xsize; x++ {
			switch a.TileAt(x, y) {
			case TREES:
				treeCount++
			case YARD:
				yardCount++
			}
		}
	}
	return yardCount * treeCount
}

func (a *Area) Step() {
	var newTiles []Tile = make(TileList, len(a.Tiles))
	for y := 0; y < a.Ysize; y++ {
		for x := 0; x < a.Xsize; x++ {
			treeCount := 0
			yardCount := 0
			var newTile Tile
			for dy := y - 1; dy <= y+1; dy++ {
				if dy < 0 || dy >= a.Ysize {
					continue
				}
				for dx := x - 1; dx <= x+1; dx++ {
					if dx < 0 || dx >= a.Xsize {
						continue
					}
					if dx == x && dy == y {
						continue
					}
					switch a.TileAt(dx, dy) {
					case TREES:
						treeCount++
					case YARD:
						yardCount++
					}
				}
			}
			switch a.TileAt(x, y) {
			case OPEN:
				if treeCount >= 3 {
					newTile = TREES
				} else {
					newTile = OPEN
				}
			case TREES:
				if yardCount >= 3 {
					newTile = YARD
				} else {
					newTile = TREES
				}
			case YARD:
				if treeCount >= 1 && yardCount >= 1 {
					newTile = YARD
				} else {
					newTile = OPEN
				}
			}
			newTiles[y*a.Xsize+x] = newTile
		}
	}
	a.Tiles = newTiles
}
