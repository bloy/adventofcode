package main

import (
	"bufio"
	"fmt"
	"strings"
)

func init() {
	AddSolution(1, solveDay1)
}

func solveDay1(pr *PuzzleRun) {
	s := bufio.NewScanner(pr.InFile)
	buf := strings.Builder{}

	for s.Scan() {
		fmt.Fprintln(&buf, s.Text())
	}
	if err := s.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	parens := strings.TrimSpace(buf.String())
	pr.ReportLoad()

	floor := strings.Count(parens, "(") - strings.Count(parens, ")")
	pr.ReportPart("Ending floor:", floor)

	floor = 0
	for i, c := range parens {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if floor == -1 {
			pr.ReportPart("First Basement:", i+1)
			break
		}
	}
}
