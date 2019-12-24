package main

import (
	"bufio"
	"fmt"
)

func init() {
	AddSolution(21, solveDay21)
}

func day21InitRobot(program string) (*Intcode, error) {
	comp, err := NewIntcodeFromInput(program)
	if err != nil {
		return nil, err
	}
	comp.AddStandardOpcodes()
	return comp, nil
}

func solveDay21(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	program := ""
	for scanner.Scan() {
		program = program + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	var script1 = `NOT A J
NOT B T
OR T J
NOT C T
OR T J
AND D J
WALK
`

	var script2 = `NOT F J
OR E J
OR H J
AND D J
NOT C T
AND T J
NOT D T
OR B T
OR E T
NOT T T
OR T J
NOT A T
OR T J
RUN
`

	pr.ReportPart(day21runpart(pr, program, script1))
	pr.ReportPart(day21runpart(pr, program, script2))
}

func day21runpart(pr *PuzzleRun, program, script string) int64 {
	in := make(chan int64, len(script))
	out := make(chan int64)
	errc := make(chan error)
	done := make(chan bool)

	var damage int64

	comp, err := day21InitRobot(program)
	pr.CheckError(err)
	comp.RunProgramChannelMode(in, out, errc, done)

	sentInput := false
	isDone := false
	for !isDone {
		select {
		case msg := <-out:
			if msg > 128 {
				damage += msg
				fmt.Println("Damage reported!")
			} else {
				fmt.Print(string(rune(msg)))
			}
		case <-done:
			isDone = true
		case err = <-errc:
			isDone = true
			fmt.Println(err)
		default:
			if !sentInput {
				sentInput = true
				for _, c := range script {
					in <- int64(c)
					fmt.Print(string(c))
				}
			}
		}
	}

	return damage
}
