package main

import "bufio"

func init() {
	AddSolution(11, solveDay11)
}

func solveDay11(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	for scanner.Scan() {
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()
}
