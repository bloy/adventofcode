package main

// Point2D is a 2 dimensional point
type Point2D struct {
	X, Y int
}

// Directional constants for a Point2D system
var (
	North = Point2D{X: 0, Y: -1}
	South = Point2D{X: 0, Y: 1}
	East  = Point2D{X: 1, Y: 0}
	West  = Point2D{X: -1, Y: 0}
)

// AbsI returns the absolute value of an int
func AbsI(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
