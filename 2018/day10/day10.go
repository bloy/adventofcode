package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const testStr string = `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>
`

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

type PointType struct {
	x, y, vx, vy int
}

func (p PointType) String() string {
	return fmt.Sprintf("(p=<%d, %d> v=<%d, %d>)", p.x, p.y, p.vx, p.vy)
}

func getInput() []PointType {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	contentStr := string(content)
	//contentStr = testStr
	points := make([]PointType, 0)
	linereg := regexp.MustCompile(`^position=<\s*(-?\d+),\s*(-?\d+)>\s*velocity=<\s*(-?\d+),\s*(-?\d+)>$`)
	for _, str := range strings.Split(contentStr, "\n") {
		match := linereg.FindStringSubmatch(str)
		if len(match) == 0 {
			continue
		}
		p := PointType{}
		p.x = atoi(match[1])
		p.y = atoi(match[2])
		p.vx = atoi(match[3])
		p.vy = atoi(match[4])
		points = append(points, p)
	}
	return points
}

func boundingBox(points []PointType) (pmin, pmax PointType) {
	if len(points) > 0 {
		pmin.x = points[0].x
		pmin.y = points[0].y
		pmax.x = points[0].x
		pmax.y = points[0].y
	}
	for _, p := range points {
		if pmin.x > p.x {
			pmin.x = p.x
		}
		if pmin.y > p.y {
			pmin.y = p.y
		}
		if pmax.x < p.x {
			pmax.x = p.x
		}
		if pmax.y < p.y {
			pmax.y = p.y
		}
	}
	return
}

func boxArea(pmin, pmax PointType) int {
	return (pmax.x - pmin.x) * (pmax.y - pmin.y)
}

func printPoints(points []PointType, time int) {
	fmt.Printf("After %d seconds:\n", time)
	pmin, pmax := boundingBox(points)
	pnormalmax := PointType{pmax.x - pmin.x, pmax.y - pmin.y, 0, 0}
	var grid [][]bool
	grid = make([][]bool, pnormalmax.y+1)
	for _, p := range points {
		x := p.x - pmin.x
		y := p.y - pmin.y
		if grid[y] == nil {
			grid[y] = make([]bool, pnormalmax.x+1)
		}
		grid[y][x] = true
	}
	for y := 0; y <= pnormalmax.y; y++ {
		for x := 0; x <= pnormalmax.x; x++ {
			switch grid[y][x] {
			case true:
				fmt.Print("#")
			case false:
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func runAll(input []PointType) {
	points := input
	pmin, pmax := boundingBox(points)
	area := boxArea(pmin, pmax)
	prevArea := area + 1
	time := 0
	for prevArea > area {
		time += 1
		for i := range points {
			points[i].x += points[i].vx
			points[i].y += points[i].vy
		}
		pmin, pmax = boundingBox(points)
		prevArea = area
		area = boxArea(pmin, pmax)
	}
	for i := range points {
		points[i].x -= points[i].vx
		points[i].y -= points[i].vy
	}
	time--
	printPoints(points, time)
}

func main() {
	input := getInput()
	runAll(input)
}
