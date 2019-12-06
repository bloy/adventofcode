package main

import (
	"bufio"
	"strings"
)

func init() {
	AddSolution(6, solveDay6)
}

type day6OrbitSpec struct {
	name   string
	parent string
}

func solveDay6(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	orbits := make(map[string]day6OrbitSpec)
	for scanner.Scan() {
		orbit := strings.Split(scanner.Text(), ")")
		orbits[orbit[1]] = day6OrbitSpec{orbit[1], orbit[0]}
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	orbits["COM"] = day6OrbitSpec{"COM", ""}
	pr.ReportLoad()

	count := 0
	for k := range orbits {
		orbiter := k
		for orbits[orbiter].parent != "" {
			count++
			orbiter = orbits[orbiter].parent
		}
	}
	pr.ReportPart(count)
}
