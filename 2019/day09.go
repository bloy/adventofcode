package main

import "bufio"

func init() {
	AddSolution(9, solveDay9)

}
func solveDay9(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	program := ""
	for scanner.Scan() {
		program = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	ic, err := NewIntcodeFromInput(program)
	if err != nil {
		pr.logger.Fatal(err)
	}
	ic.AddStandardOpcodes()

	out, err := ic.RunProgram([]int64{1})
	if err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportPart(out)

	ic, err = NewIntcodeFromInput(program)
	if err != nil {
		pr.logger.Fatal(err)
	}
	ic.AddStandardOpcodes()

	out, err = ic.RunProgram([]int64{2})
	if err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportPart(out)
}
