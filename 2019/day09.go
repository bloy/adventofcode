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
	pr.logger.Println(program)

}
