package main

import (
	"fmt"
)

type Point struct {
	X, Y, Z int
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func (p Point) String() string {
	return fmt.Sprintf("<%d, %d, %d>", p.X, p.Y, p.Z)
}

func (p1 Point) Distance(p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y) + abs(p1.Z-p2.Z)
}

type Bot struct {
	Position Point
	Radius   int
}
