package main

import "fmt"

func main() {
	depth := inputDepth
	target := inputTarget
	fmt.Println(depth, target)
	cave := NewCave(depth, target.X, target.Y)
	fmt.Println(cave)
	totalRisk := 0
	for y := 0; y <= target.Y; y++ {
		for x := 0; x <= target.X; x++ {
			r := cave.Region(x, y)
			totalRisk += r.Risk()
		}
	}
	fmt.Println("part 1 total risk:", totalRisk)

}
