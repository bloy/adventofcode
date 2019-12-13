package main

import "bufio"

func init() {
	AddSolution(13, solveDay13)
}

func solveDay13(pr *PuzzleRun) {
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
}
