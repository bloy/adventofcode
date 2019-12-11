package main

import (
	"bufio"
	"fmt"
)

func init() {
	AddSolution(11, solveDay11)
}

func solveDay11(pr *PuzzleRun) {
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
	if err != nil {
		pr.logger.Fatal(err)
	}

	grid := NewGrid()
	turns := map[Point]map[int64]Point{
		Up:    map[int64]Point{0: Left, 1: Right},
		Down:  map[int64]Point{0: Right, 1: Left},
		Left:  map[int64]Point{0: Down, 1: Up},
		Right: map[int64]Point{0: Up, 1: Down},
	}
	dir := Up
	pos := Point{0, 0}

	comp.AddStandardOpcodes()

	robotInputOpcode := func(ic *Intcode, positions []int64) (done bool, err error) {
		if ic.Verbose {
			fmt.Println("RIN ", positions)
		}
		c := grid.GetPoint(pos)
		if ic.Verbose {
			fmt.Println(string(c), "at position", pos)
		}
		o := positions[0]
		if c == '#' {
			ic.mem[o] = 1
		} else {
			ic.mem[o] = 0
		}
		ic.pc += 2
		return
	}

	expectColorOutput := true
	robotOutputOpcode := func(ic *Intcode, positions []int64) (done bool, err error) {
		if ic.Verbose {
			fmt.Println("ROUT", positions, expectColorOutput)
		}
		in := positions[0]
		if expectColorOutput {
			expectColorOutput = false
			c := '.'
			if in == 1 {
				c = '#'
			}
			grid.SetPoint(pos, c)
		} else {
			expectColorOutput = true
			dir = turns[dir][in]
			pos = Point{X: pos.X + dir.X, Y: pos.Y + dir.Y}
		}
		ic.pc += 2
		return
	}

	comp.AddOpcode(3, 1, "w", robotInputOpcode)
	comp.AddOpcode(4, 1, "r", robotOutputOpcode)
	//comp.Verbose = true
	_, err = comp.RunProgram(nil)
	count := len(grid.values)
	pr.ReportPart(count)

	// -----------------------------------------------------------

	comp, err = NewIntcodeFromInput(program)
	if err != nil {
		pr.logger.Fatal(err)
	}

	grid = NewGrid()
	dir = Up
	pos = Point{0, 0}

	comp.AddStandardOpcodes()
	comp.AddOpcode(3, 1, "w", robotInputOpcode)
	comp.AddOpcode(4, 1, "r", robotOutputOpcode)

	grid.SetPoint(pos, '#')
	_, err = comp.RunProgram(nil)
	pr.ReportPart("\n" + grid.String())
}
