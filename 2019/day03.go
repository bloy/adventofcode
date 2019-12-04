package main

import (
	"bufio"
	"strconv"
	"strings"
)

func init() {
	AddSolution(3, solveDay3)
}

type day3pathSegment struct {
	direction Point
	distance  int
}

func solveDay3(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	wires := make([][]day3pathSegment, 0)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ",")
		wire := make([]day3pathSegment, len(strs))
		for i, str := range strs {
			seg := day3pathSegment{}
			switch str[:1] {
			case "U":
				seg.direction = Up
			case "D":
				seg.direction = Down
			case "L":
				seg.direction = Left
			case "R":
				seg.direction = Right
			}
			num, err := strconv.Atoi(str[1:])
			if err != nil {
				pr.logger.Fatal(err)
			}
			seg.distance = num
			wire[i] = seg
		}
		wires = append(wires, wire)
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	grid := NewGrid()
	origin := Point{0, 0}
	grid.SetPoint(origin, 'o')
	crossings := make([]Point, 0)
	for i, wire := range wires {
		current := origin
		wirechar := '&'
		if i == 1 {
			wirechar = '*'
		}
		for _, seg := range wire {
			for i := 0; i < seg.distance; i++ {
				current = current.Add(seg.direction)
				if grid.GetPoint(current) == '&' && wirechar == '*' {
					grid.SetPoint(current, 'X')
					crossings = append(crossings, current)
				} else {
					grid.SetPoint(current, wirechar)
				}
			}
		}
	}

	min := -1
	for _, c := range crossings {
		dist := origin.Distance(c)
		if min == -1 || dist < min {
			min = dist
		}
	}
	pr.ReportPart("Closest Crossing:", min)

	min = -1
	for _, crossing := range crossings {
		wsteps := make([]int, 2)
		for i, wire := range wires {
			wsteps[i] = day3stepsTo(wire, crossing)
		}
		total := wsteps[0] + wsteps[1]
		if min == -1 || min > total {
			min = total
		}
	}
	pr.ReportPart("Minimum wire steps to crossing:", min)
}

func day3stepsTo(wire []day3pathSegment, p Point) int {
	steps := 0
	current := Point{0, 0}
	for _, seg := range wire {
		for i := 0; i < seg.distance; i++ {
			current = current.Add(seg.direction)
			steps++
			if current.X == p.X && current.Y == p.Y {
				return steps
			}
		}
	}
	return -1
}
