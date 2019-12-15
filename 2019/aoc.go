package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// SolutionFunc is a function over a specific puzzle
type SolutionFunc func(*PuzzleRun)

// Solution gathers data for a single puzzle solution
type Solution struct {
	Parts []SolutionFunc
}

// PuzzleRun structs hold data on a running puzzle
type PuzzleRun struct {
	InFile    *os.File
	needClose bool
	logger    *log.Logger
	timings   struct {
		start time.Time
		load  time.Time
		parts []time.Time
		done  time.Time
	}
}

func newPuzzleRun(n int) (*PuzzleRun, error) {
	pr := &PuzzleRun{
		logger: log.New(os.Stdout, "", 0),
	}
	name := fmt.Sprintf("day%02d.txt", n)
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	pr.InFile = f
	pr.needClose = true
	return pr, nil
}

// Close closes any open resources
func (pr *PuzzleRun) Close() {
	if pr.needClose {
		pr.InFile.Close()
		pr.needClose = false
	}
}

// CheckError checks for a nil error
func (pr *PuzzleRun) CheckError(err error) {
	if err != nil {
		pr.logger.Fatal(err)
	}
}

// ReportLoad stores the loading time for the report
func (pr *PuzzleRun) ReportLoad() {
	pr.timings.load = time.Now()
	pr.logger.Println("Loaded data")
}

// ReportPart is a method on *PuzzleRun
func (pr *PuzzleRun) ReportPart(v ...interface{}) {
	t := time.Now()
	pr.timings.parts = append(pr.timings.parts, t)
	partNum := len(pr.timings.parts)
	tokens := append([]interface{}{fmt.Sprintf("Part %d:", partNum)}, v...)
	pr.logger.Println(tokens...)
}

func (pr *PuzzleRun) printTimings() {
	prev := pr.timings.start
	pr.logger.Println("Timings:")
	if !pr.timings.load.IsZero() {
		pr.logger.Println("    Load:", pr.timings.load.Sub(prev))
		prev = pr.timings.load
	}
	for i, part := range pr.timings.parts {
		pr.logger.Printf("    Part %d: %v\n", i+1, part.Sub(prev))
		prev = part
	}
	pr.logger.Println("    Total:", pr.timings.done.Sub(pr.timings.start))
}

var solutions = make(map[int][]Solution)

// AddSolution adds the given list of parts as a solution to a problem
func AddSolution(problem int, fns ...SolutionFunc) {
	solutions[problem] = append(solutions[problem], Solution{fns})
}

// FindSolution finds the given solution
func FindSolution(problem, solutionNumber int) (*Solution, error) {
	sList, ok := solutions[problem]
	if !ok {
		return nil, fmt.Errorf("No solutions for problem %d", problem)
	}
	if solutionNumber > len(sList) {
		return nil, fmt.Errorf("Problem %d has only %d solutions", problem, len(sList))
	}
	return &sList[solutionNumber-1], nil
}
