package main

import (
	"container/list"
	"fmt"
)

func canFlowSideways(area *Area, p Point) bool {
	down := p.Down()
	return area.Tiles[p] != CLAY && (area.Tiles[down] == SETTLED || area.Tiles[down] == CLAY)
}

func doAll(area *Area, source Point) {
	todo := list.New()
	firstPoint := source
	firstPoint.y = area.MinPoint.y
	todo.PushBack(firstPoint)
	area.Tiles[firstPoint] = FLOWING
	for todo.Front() != nil {
		p := todo.Front().Value.(Point)
		todo.Remove(todo.Front())
		if area.Tiles[p] != FLOWING {
			continue
		}
		if p.y < area.MinPoint.y || p.y > area.MaxPoint.y {
			continue
		}
		down := p.Down()
		if area.Tiles[down] == SAND {
			area.Tiles[down] = FLOWING
			todo.PushBack(down)
		} else if area.Tiles[down] == CLAY || area.Tiles[down] == SETTLED {
			todo.PushBack(p.Up())
			var enclosedLeft bool = false
			var enclosedRight bool = false
			flow := p
			for canFlowSideways(area, flow) {
				flow = flow.Left()
			}
			if area.Tiles[flow] == CLAY {
				enclosedLeft = true
			}
			flow = p
			for canFlowSideways(area, flow) {
				flow = flow.Right()
			}
			if area.Tiles[flow] == CLAY {
				enclosedRight = true
			}
			flow = p
			for canFlowSideways(area, flow) {
				if enclosedLeft && enclosedRight {
					todo.PushBack(flow.Up())
					area.Tiles[flow] = SETTLED
				} else {
					area.Tiles[flow] = FLOWING
				}
				flow = flow.Left()
			}
			if area.Tiles[flow] != CLAY && area.Tiles[flow.Down()] == SAND {
				area.Tiles[flow] = FLOWING
				todo.PushBack(flow)
			}
			flow = p
			for canFlowSideways(area, flow) {
				if enclosedLeft && enclosedRight {
					todo.PushBack(flow.Up())
					area.Tiles[flow] = SETTLED
				} else {
					area.Tiles[flow] = FLOWING
				}
				flow = flow.Right()
			}
			down = flow.Down()
			if area.Tiles[flow] != CLAY && area.Tiles[flow.Down()] == SAND {
				area.Tiles[flow] = FLOWING
				todo.PushBack(flow)
			}
		}
	}
	var waterCount int
	var settledCount int
	for p, tile := range area.Tiles {
		if p.y < area.MinPoint.y || p.y > area.MaxPoint.y {
			continue
		}
		if tile == FLOWING || tile == SETTLED {
			waterCount++
			if tile == SETTLED {
				settledCount++
			}
		}
	}
	fmt.Println("Part 1 - Number of water tiles:", waterCount)
	fmt.Println("Part 2 - Number of settled water tiles: ", settledCount)
}

func main() {
	area := ParseInput(inputData)
	source := Point{500, 0}
	doAll(area, source)
}
