package main

import (
	"bufio"
	"strconv"
)

func init() {
	AddSolution(1, solveDay1)
}

func solveDay1(pr *PuzzleRun) {

	scanner := bufio.NewScanner(pr.InFile)
	masses := make([]int, 0)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			pr.logger.Fatal(err)
		}
		masses = append(masses, num)
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	fuel := 0
	for _, mass := range masses {
		fuel += mass/3 - 2
	}

	pr.ReportPart("Fuel Required", fuel)

	fuel = 0
	for _, mass := range masses {
		newMass := mass
		newFuel := newMass/3 - 2
		for newFuel > 0 {
			fuel += newFuel
			newMass = newFuel
			newFuel = newMass/3 - 2
		}
	}
	pr.ReportPart("Fuel Required", fuel)
}
