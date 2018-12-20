package main

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

func NewRoom(x, y int) *Room {
	var r *Room = &Room{}
	r.Doors = make(map[Direction]*Room)
	r.Point = Point{x, y}
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
	b.StartRoom = NewRoom(0, 0)
	b.Rooms[b.StartRoom.Position] = b.StartRoom
}
