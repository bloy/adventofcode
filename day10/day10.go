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
	contentStr = testStr
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

func runPart1(input []PointType) {
	fmt.Println("part 1", input)
}

func runPart2(input []PointType) {
	fmt.Println("part 2")
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
