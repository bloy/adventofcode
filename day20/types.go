package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type Room struct {
	Doors    map[Direction]*Room
	Position Point
}

func NewRoom(p Point) *Room {
	var r *Room = &Room{}
	r.Doors = make(map[Direction]*Room)
	r.Position = Point{p.X, p.Y}
	return r
}

type Base struct {
	Rooms     map[Point]*Room
	StartRoom *Room
	MinPos    Point
	MaxPos    Point
}

func NewBase() *Base {
	var b *Base = &Base{}
	b.Rooms = make(map[Point]*Room)
	b.MinPos = Point{0, 0}
	b.MaxPos = Point{0, 0}
	b.StartRoom = NewRoom(Point{0, 0})
	b.Rooms[b.StartRoom.Position] = b.StartRoom
	return b
}

func (b *Base) AddRoom(dir Direction, fromRoom *Room) (toRoom *Room) {
	newPoint := Point{fromRoom.Position.X, fromRoom.Position.Y}
	var reverseDir Direction
	switch dir {
	case NORTH:
		reverseDir = SOUTH
		newPoint.Y = newPoint.Y - 1
	case EAST:
		reverseDir = WEST
		newPoint.X = newPoint.X + 1
	case SOUTH:
		reverseDir = NORTH
		newPoint.Y = newPoint.Y + 1
	case WEST:
		reverseDir = EAST
		newPoint.X = newPoint.X - 1
	}
	if newPoint.X < b.MinPos.X {
		b.MinPos.X = newPoint.X
	}
	if newPoint.X > b.MaxPos.X {
		b.MaxPos.X = newPoint.X
	}
	if newPoint.Y < b.MinPos.Y {
		b.MinPos.Y = newPoint.Y
	}
	if newPoint.Y > b.MaxPos.Y {
		b.MaxPos.Y = newPoint.Y
	}
	toRoom = b.Rooms[newPoint]
	if toRoom == nil {
		fmt.Println("new room at", newPoint)
		toRoom = NewRoom(newPoint)
		b.Rooms[newPoint] = toRoom
	}
	fromRoom.Doors[dir] = toRoom
	toRoom.Doors[reverseDir] = fromRoom
	return toRoom
}
