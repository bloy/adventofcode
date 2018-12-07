package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const testStr string = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

type Point struct {
	x, y int
}

type Box struct {
	xmin, xmax, ymin, ymax int
}

func getInput() []Point {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(content)
	//str = testStr
	str = strings.TrimSpace(str)
	pointstrs := strings.Split(str, "\n")
	points := make([]Point, 0, len(pointstrs))
	for _, s := range pointstrs {
		values := strings.Split(s, ", ")
		x, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		p := Point{x, y}
		points = append(points, p)
	}
	return points
}

func boundingBox(points []Point, extra int) Box {
	var b Box
	if len(points) != 0 {
		b.xmin = points[0].x
		b.xmax = points[0].x
		b.ymin = points[0].y
		b.ymax = points[0].y
	}
	for _, p := range points {
		if p.x < b.xmin {
			b.xmin = p.x
		}
		if p.x > b.xmax {
			b.xmax = p.x
		}
		if p.y < b.ymin {
			b.ymin = p.y
		}
		if p.y > b.ymax {
			b.ymax = p.y
		}
	}
	b.xmin -= extra
	b.ymin -= extra
	b.xmax += extra
	b.ymax += extra
	return b
}

func (b Box) boxPoints() []Point {
	points := make([]Point, 0, (b.xmax-b.xmin)*(b.ymax-b.ymin))
	for y := b.ymin; y <= b.ymax; y++ {
		for x := b.xmin; x <= b.xmax; x++ {
			points = append(points, Point{x, y})
		}
	}
	return points
}

func distance(p1, p2 Point) int {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	if dx < 0 {
		dx = -1 * dx
	}
	if dy < 0 {
		dy = -1 * dy
	}
	return dx + dy
}

func closest(p Point, points []Point) int {
	minDist := 1000
	minPoint := -1
	for i, dp := range points {
		dist := distance(p, dp)
		if dist < minDist {
			minPoint = i
			minDist = dist
		} else if dist == minDist {
			minPoint = -1
			minDist = dist
		}
	}
	return minPoint
}

func runPart1(input []Point) {
	bounds := boundingBox(input, 0)
	pointCounts := make(map[int]int)
	invalidPoints := make(map[int]bool)
	for _, p := range bounds.boxPoints() {
		closePoint := closest(p, input)
		if closePoint >= 0 {
			pointCounts[closePoint] = pointCounts[closePoint] + 1
			if p.x == bounds.xmin || p.x == bounds.xmax || p.y == bounds.ymin || p.y == bounds.ymax {
				invalidPoints[closePoint] = true
			}
		}
	}
	maxCount := 0
	for i, count := range pointCounts {
		if invalidPoints[i] == true {
			continue
		}
		if count > maxCount {
			maxCount = count
		}
	}
	fmt.Println("part 1", maxCount)
}

func runPart2(input []Point) {
	const regionDistance int = 10000
	//const regionDistance int = 32
	pointq := make([]Point, 0, len(input)*5)
	seenPoints := make(map[Point]bool)
	regionCount := 0
	for _, p := range input {
		//fmt.Println(p)
		for y := p.y - 1; y <= p.y+1; y++ {
			for x := p.x - 1; x <= p.x+1; x++ {
				newp := Point{x, y}
				_, ok := seenPoints[newp]
				if !ok {
					pointq = append(pointq, newp)
					seenPoints[newp] = true
				}
			}
		}
	}
	// queue primed
	qCount := 0
	for len(pointq) > 0 {
		// while there's items on the queue
		p := pointq[0]
		pointq = pointq[1:] // pop an item off the queue
		qCount += 1
		dist := 0
		for _, pt := range input {
			dist += distance(p, pt)
		}
		//fmt.Printf("%d, %d ", p, dist)
		if dist < regionDistance {
			regionCount++
			//fmt.Printf("! %d", regionCount)
			for y := p.y - 1; y <= p.y+1; y++ {
				for x := p.x - 1; x <= p.x+1; x++ {
					var newp Point = Point{x, y}
					_, ok := seenPoints[newp]
					if !ok {
						pointq = append(pointq, newp)
						seenPoints[newp] = true
					}
				}
			}
		}
		//fmt.Print("\n")
	}
	fmt.Println("part 2", regionCount)
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
