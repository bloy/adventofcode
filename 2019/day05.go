package main

import "bufio"

func init() {
	AddSolution(5, solveDay5)
}

func solveDay5(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	var ic *Intcode
	var err error
	for scanner.Scan() {
		ic, err = NewIntcodeFromInput(scanner.Text())
		if err != nil {
			pr.logger.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	ic.AddStandardOpcodes()
	pr.ReportLoad()
	outputs, err := ic.RunProgram([]int{1})
	pr.ReportPart("Part1", outputs, err)
}
