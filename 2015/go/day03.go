package main

import (
	"bufio"
	"fmt"
	"strings"
)

func init() {
	AddSolution(3, solveDay3)
}

func solveDay3(pr *PuzzleRun) {
	s := bufio.NewScanner(pr.InFile)
	buf := strings.Builder{}
	for s.Scan() {
		fmt.Fprint(&buf, s.Text())
	}
	if err := s.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	dirs := buf.String()
	pr.ReportLoad()

	directions := map[rune]Point{
		'^': Point{X: 0, Y: -1},
		'v': Point{X: 0, Y: 1},
		'<': Point{X: -1, Y: 0},
		'>': Point{X: 1, Y: 0},
	}

	santa := Point{0, 0}
	houses := make(map[Point]int)
	houses[santa]++
	for _, c := range dirs {
		dir := directions[c]
		santa.X += dir.X
		santa.Y += dir.Y
		houses[santa]++
	}

	pr.ReportPart(len(houses))

	santa = Point{0, 0}
	robo := Point{0, 0}
	houses = make(map[Point]int)
	houses[santa] = 2
	for i, c := range dirs {
		dir := directions[c]
		if i%2 == 0 {
			santa.X += dir.X
			santa.Y += dir.Y
			houses[santa]++
		} else {
			robo.X += dir.X
			robo.Y += dir.Y
			houses[robo]++
		}
	}
	pr.ReportPart(len(houses))
}
