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

	pr.ReportPart(day21part1(pr, program))
}

var day21part1script = `NOT A J
NOT B T
OR T J
NOT C T
OR T J
AND D J
WALK
`

func day21part1(pr *PuzzleRun, program string) int64 {
	in := make(chan int64, len(day21part1script))
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
				for _, c := range day21part1script {
					in <- int64(c)
					fmt.Print(string(c))
				}
			}
		}
	}

	return damage
}
