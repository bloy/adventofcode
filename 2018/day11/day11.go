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
	var subGrids []int = make([]int, gridSize*gridSize)
	var maxPower int
	var maxPowerPoint Point
	var maxPowerSize int
	for x := 1; x <= gridSize; x++ {
		rackID := x + 10
		columnTotal := 0
		for y := 1; y <= gridSize; y++ {
			adjacentBoxTotal := 0
			if x > 1 {
				adjacentBoxTotal = subGrids[(y-1)*gridSize+(x-2)]
			}
			level := ((rackID * y) + serial) * rackID
			level = level / 100
			level = level % 10
			level = level - 5
			columnTotal += level
			subGrids[(y-1)*gridSize+(x-1)] = adjacentBoxTotal + columnTotal
		}
	}

	// TODO: precompute 1,1-x,y subgrid values for all subgrids
	// subgrid calc is summed area(x2,y2) - summed area(x2,y1) - summed area(x1,y2) + summed area(x1,y1)
	for size := 1; size <= gridSize; size++ {
		for y := 1; y <= gridSize-size+1; y++ {
			for x := 1; x <= gridSize-size+1; x += 1 {
				// quadrants:
				//   Q1 Q2
				//   Q3 Q4
				// want the Q4 total
				// have the total of all, Q1, Q2 + Q1 (Q12), and Q3 + Q1 (Q13)
				// Q4 = ALL - Q13 - Q12 + Q1
				x1 := x - 2
				x2 := x - 1 + size - 1
				y1 := y - 2
				y2 := y - 1 + size - 1
				q1Total := 0
				q12Total := 0
				q13Total := 0
				allTotal := subGrids[y2*gridSize+x2]
				if x1 >= 0 && y1 >= 0 {
					q1Total = subGrids[y1*gridSize+x1]
					q12Total = subGrids[y1*gridSize+x2]
					q13Total = subGrids[y2*gridSize+x1]
				} else if x1 >= 0 {
					q13Total = subGrids[y2*gridSize+x1]
				} else if y1 >= 0 {
					q12Total = subGrids[y1*gridSize+x2]
				}
				power := allTotal - q12Total - q13Total + q1Total
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
