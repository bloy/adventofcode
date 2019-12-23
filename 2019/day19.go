package main

import (
	"bufio"
	"fmt"
)

func init() {
	AddSolution(19, solveDay19)
}

func solveDay19(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	program := ""
	for scanner.Scan() {
		program = program + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	grid := NewGrid()

	getPoint := func(p Point) rune {
		if r, ok := grid.values[p]; ok {
			return r
		}
		comp, err := NewIntcodeFromInput(program)
		pr.CheckError(err)
		comp.AddStandardOpcodes()
		out, err := comp.RunProgram([]int64{int64(p.X), int64(p.Y)})
		v := out[0]
		c := '?'
		if v == 1 {
			c = '#'
		} else if v == 0 {
			c = '.'
		}
		grid.SetPoint(p, c)
		return c
	}

	count := 0
	size := 50
	fmt.Print("   ")
	for x := 0; x < size; x++ {
		fmt.Print(x % 10)
	}
	fmt.Print("\n")
	for y := 0; y < size; y++ {
		fmt.Printf("%2d ", y)
		for x := 0; x < size; x++ {
			c := getPoint(Point{X: x, Y: y})
			if c == '#' {
				count++
			}
			fmt.Print(string(c))
		}
		fmt.Print("\n")
	}

	pr.ReportPart(count)
}
