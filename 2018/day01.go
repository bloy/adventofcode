package main

import (
	"bufio"
	"strconv"
)

var day1Data []int

func init() {
	AddSolution(1, day1part1, day1part2)
}

func day1Load(pr *PuzzleRun) []int {
	if day1Data != nil {
		return day1Data
	}
	s := bufio.NewScanner(pr.InFile)
	for s.Scan() {
		num, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			pr.logger.Fatalln("Numeric parsing error:", err)
		}
		day1Data = append(day1Data, int(num))
	}
	pr.ReportLoad()
	return day1Data
}

func day1part1(pr *PuzzleRun) {
	changes := day1Load(pr)
	freq := 0
	for _, change := range changes {
		freq += change
	}
	pr.ReportPart(freq)
}

func day1part2(pr *PuzzleRun) {
	changes := day1Load(pr)
	seen := make(map[int]bool)
	seen[0] = true
	var steps, cycles, freq int
	for true {
		for _, change := range changes {
			steps++
			freq += change
			if seen[freq] == true {
				pr.ReportPart("Frequency:", freq, "Steps:", steps, "Cycles:", cycles)
				return
			}
			seen[freq] = true
		}
		cycles++
	}
}
