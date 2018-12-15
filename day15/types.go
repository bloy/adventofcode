package main

import (
	"fmt"
	"strings"
)

type Square byte

const (
	Space Square = iota
	Wall
)

func (s Square) String() string {
	switch s {
	case Space:
		return "."
	case Wall:
		return "#"
	default:
		return " "
	}
}

type Level struct {
	Xsize   int
	Ysize   int
	squares []Square
}

func NewLevel(str string) (level *Level) {
	str = strings.TrimSpace(str)
	lines := strings.Split(str, "\n")
	level = &Level{}
	level.Ysize = len(lines)
	level.Xsize = len(lines[0])
	level.squares = make([]Square, level.Xsize*level.Ysize)
	for y, line := range lines {
		for x, ch := range line {
			square := Space
			if ch == '#' {
				square = Wall
			}
			level.squares[y*level.Xsize+x] = square
		}
	}
	return level
}

func (level *Level) String() string {
	var out strings.Builder
	for y := 0; y < level.Ysize; y++ {
		for x := 0; x < level.Xsize; x++ {
			fmt.Fprintf(&out, "%v", level.squares[y*level.Xsize+x])
		}
		fmt.Fprint(&out, "\n")
	}
	return out.String()
}
