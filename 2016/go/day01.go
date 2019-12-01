package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	AddSolution(1, day1)
}

func day1(pr *PuzzleRun) {
	type instruction struct {
		dir string
		amt int
	}

	facingMachine := map[Point2D]map[string]Point2D{
		North: map[string]Point2D{"L": West, "R": East},
		South: map[string]Point2D{"L": East, "R": West},
		East:  map[string]Point2D{"L": North, "R": South},
		West:  map[string]Point2D{"L": South, "R": North},
	}

	inbytes, err := ioutil.ReadAll(pr.InFile)
	if err != nil {
		pr.logger.Fatalln(err)
	}
	iList := make([]instruction, 0)
	for _, str := range strings.Split(string(inbytes), ", ") {
		dir := str[:1]
		num, err := strconv.Atoi(strings.TrimSpace(str[1:]))
		if err != nil {
			pr.logger.Fatalln(err)
		}
		iList = append(iList, instruction{dir: dir, amt: num})
	}
	pr.ReportLoad()

	current := Point2D{0, 0}
	facing := North
	for _, instr := range iList {
		facing = facingMachine[facing][instr.dir]
		for i := 0; i < instr.amt; i++ {
			current.X += facing.X
			current.Y += facing.Y
		}
	}
	dist := AbsI(current.X) + AbsI(current.Y)
	pr.ReportPart("Distance:", dist)

	current = Point2D{0, 0}
	facing = North
	seen := make(map[Point2D]bool)
	seen[current] = true
	for _, instr := range iList {
		facing = facingMachine[facing][instr.dir]
		for i := 0; i < instr.amt; i++ {
			current.X += facing.X
			current.Y += facing.Y
			if seen[current] {
				dist := AbsI(current.X) + AbsI(current.Y)
				pr.ReportPart("Distance:", dist)
				return
			}
			seen[current] = true
		}
	}
}
