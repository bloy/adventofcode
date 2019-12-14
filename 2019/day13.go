package main

import (
	"bufio"
	"fmt"
	"time"
)

func init() {
	AddSolution(13, solveDay13)
}

func solveDay13(pr *PuzzleRun) {
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
	if err != nil {
		pr.logger.Fatal(err)
	}
	comp.AddStandardOpcodes()
	tiles := []rune{' ', 0x2588, 0x2591, 0x2500, 0x2022}
	outputs, err := comp.RunProgram(nil)
	if err != nil {
		pr.logger.Fatal(err)
	}
	grid := NewGrid()
	for i := 0; i < len(outputs); i += 3 {
		x := outputs[i]
		y := outputs[i+1]
		tile := tiles[outputs[i+2]]
		grid.SetPoint(Point{int(x), int(y)}, tile)
	}
	pr.logger.Print(grid)
	pr.logger.Println()
	minPoint, maxPoint := grid.Bounds()
	count := 0
	for y := minPoint.Y; y <= maxPoint.Y; y++ {
		for x := minPoint.X; x <= maxPoint.X; x++ {
			t := grid.GetPoint(Point{x, y})
			if t == tiles[2] {
				count++
			}
		}
	}
	pr.ReportPart(count)

	comp, err = NewIntcodeFromInput(program)
	if err != nil {
		pr.logger.Fatal(err)
	}
	comp.AddStandardOpcodes()

	grid = NewGrid()

	var outputState, blocks int
	var score, paddleX, ballX int64
	outputs = make([]int64, 3)
	gameOutputOpcode := func(ic *Intcode, positions []int64) (done bool, err error) {
		in := positions[0]
		outputs[outputState] = in
		outputState = (outputState + 1) % 3
		if outputState == 0 {
			if outputs[0] == -1 && outputs[1] == 0 {
				score = outputs[2]
			} else {
				tile := tiles[outputs[2]]
				grid.SetPoint(Point{int(outputs[0]), int(outputs[1])}, tile)
				if outputs[2] == 4 { // ball position
					ballX = outputs[0]
				} else if outputs[2] == 3 { // paddle position
					paddleX = outputs[0]
				}
				min, max := grid.Bounds()
				blocks = 0
				for y := min.Y; y <= max.Y; y++ {
					for x := min.X; x <= max.Y; x++ {
						if grid.GetPoint(Point{X: x, Y: y}) == tiles[2] {
							blocks++
						}
					}
				}
			}
		}
		ic.pc += 2
		return
	}

	gameInputOpcode := func(ic *Intcode, positions []int64) (done bool, err error) {
		fmt.Print("\x1b[2;1H")
		fmt.Println(grid)
		fmt.Printf("Score: %d     Blocks: %d", score, blocks)
		time.Sleep(time.Millisecond * 10)
		o := positions[0]
		if ballX < paddleX {
			ic.mem[o] = -1
		} else if ballX > paddleX {
			ic.mem[o] = 1
		} else {
			ic.mem[o] = 0
		}
		ic.pc += 2
		return
	}
	fmt.Print("\x1b[s")
	comp.AddOpcode(3, 1, "w", gameInputOpcode)
	comp.AddOpcode(4, 1, "r", gameOutputOpcode)
	comp.mem[0] = 2
	_, err = comp.RunProgram(nil)
	if err != nil {
		pr.logger.Fatal(err)
	}
	fmt.Print("\x1b[2;1H")
	fmt.Println(grid)
	fmt.Printf("Score: %d     Blocks: %d", score, blocks)
	fmt.Print("\x1b[u")
	pr.ReportPart(score)
}
