package main

import "bufio"

func init() {
	AddSolution(5, solveDay5)
}

func solveDay5(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	ic, err := NewIntcodeFromInput(program)
	if err != nil {
		pr.logger.Fatal(err)
	}
	ic.AddStandardOpcodes()
	pr.ReportLoad()
	outputs, err := ic.RunProgram([]int{1})
	pr.ReportPart("Part1", outputs, err)

	ic, err = NewIntcodeFromInput(program)
	if err != nil {
		pr.logger.Fatal(err)
	}
	ic.AddStandardOpcodes()
	outputs, err = ic.RunProgram([]int{5})
	pr.ReportPart("Part2", outputs, err)
}
