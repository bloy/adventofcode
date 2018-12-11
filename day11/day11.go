package main

import (
	"fmt"
)

const input = 7347

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func greatestPower1(serial, gridSize int) {
	var cells map[Point]int = make(map[Point]int, gridSize*gridSize)
	var maxPower int = -9999
	var maxPowerPoint Point
	for y := 1; y <= gridSize-3; y++ {
		for x := 1; x <= gridSize-3; x++ {
			power := 0
			for suby := y; suby < y+3; suby++ {
				for subx := x; subx < x+3; subx++ {
					subcell := Point{subx, suby}
					if value, ok := cells[subcell]; ok == true {
						power += value
					} else {
						rackID := subcell.x + 10
						level := ((rackID * subcell.y) + serial) * rackID
						level = level / 100
						level = level % 10
						level = level - 5
						cells[subcell] = level
						power += level
					}
				}
			}
			if power > maxPower {
				maxPower = power
				maxPowerPoint = Point{x, y}
			}
		}
	}
	fmt.Printf("serial number %d: %v\n", serial, maxPowerPoint)
}

func greatestPower2(serial, gridSize int) {
	var cells map[Point]int = make(map[Point]int, gridSize*gridSize)
	var maxPower int = -9999
	var maxPowerPoint Point
	var maxPowerSize int
	// TODO: precompute cell values
	// TODO: precompute 1,1-x,y subgrid values for all subgrids
	// subgrid calc is summed area(x2,y2) - summed area(x2,y1) - summed area(x1,y2) + summed area(x1,y1)
	for size := 1; size <= gridSize; size++ {
		fmt.Println(size)
		for y := 1; y <= gridSize-size+1; y++ {
			for x := 1; x <= gridSize-size+1; x += 1 {
				power := 0
				for suby := y; suby < y+size; suby++ {
					for subx := x; subx < x+size; subx++ {
						subcell := Point{subx, suby}
						if value, ok := cells[subcell]; ok == true {
							power += value
						} else {
							rackID := subcell.x + 10
							level := ((rackID * subcell.y) + serial) * rackID
							level = level / 100
							level = level % 10
							level = level - 5
							cells[subcell] = level
							power += level
						}
					}
				}
				if power > maxPower {
					maxPower = power
					maxPowerPoint = Point{x, y}
					maxPowerSize = size
				}
			}
		}
	}
	fmt.Printf("serial number %d: %v,%d\n", serial, maxPowerPoint, maxPowerSize)
}

func main() {
	greatestPower1(18, 300)
	greatestPower1(42, 300)
	greatestPower1(input, 300)

	greatestPower2(18, 300)
	greatestPower2(42, 300)
	greatestPower2(input, 300)
}
