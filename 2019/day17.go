package main

import (
	"bufio"
	"fmt"
)

func init() {
	AddSolution(17, solveDay17)
}

func solveDay17(pr *PuzzleRun) {
	fmt.Print("\x1b[2J\x1b[H")
	scanner := bufio.NewScanner(pr.InFile)
	program := ""
	for scanner.Scan() {
		program = program + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	comp, err := NewIntcodeFromInput(program)
	pr.CheckError(err)
	comp.AddStandardOpcodes()
	outputs, err := comp.RunProgram(nil)
	grid := NewGrid()
	cursor := Point{0, 0}
	for _, n := range outputs {
		if n == 10 {
			cursor = Point{X: 0, Y: cursor.Y + 1}
			continue
		}
		grid.SetPoint(cursor, rune(n))
		cursor = Point{X: cursor.X + 1, Y: cursor.Y}
	}
	fmt.Println(grid)
	min, max := grid.Bounds()
	sum := 0
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			point := Point{X: x, Y: y}
			if grid.GetPoint(point) != '#' {
				continue
			}
			if grid.GetPoint(point.Add(North)) != '#' {
				continue
			}
			if grid.GetPoint(point.Add(South)) != '#' {
				continue
			}
			if grid.GetPoint(point.Add(East)) != '#' {
				continue
			}
			if grid.GetPoint(point.Add(West)) != '#' {
				continue
			}
			sum += point.X * point.Y
		}
	}
	pr.ReportPart(sum)
}
