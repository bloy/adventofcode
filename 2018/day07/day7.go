package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type depType map[string]string

const testStr string = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`

func getInput() depType {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(content)
	//str = testStr
	strs := strings.Split(strings.TrimSpace(str), "\n")
	deps := make(depType)
	for _, str := range strs {
		parts := strings.Split(str, " ")
		after := parts[1]
		before := parts[7]
		if _, ok := deps[after]; !ok {
			deps[after] = ""
		}
		deps[before] = deps[before] + after
	}
	return deps
}

func getNextStep(stepMap depType) string {
	var next string
	for step, deps := range stepMap {
		if len(deps) == 0 && (step < next || next == "") {
			next = step
		}
	}
	return next
}

func startStep(stepMap depType, doStep string) depType {
	var newDeps depType = make(depType)
	for step, deps := range stepMap {
		if step == doStep {
			continue
		}
		newDeps[step] = deps
	}
	return newDeps
}

func finishStep(stepMap depType, doStep string) depType {
	var newDeps depType = make(depType)
	for step, deps := range stepMap {
		if step == doStep {
			continue
		}
		newDeps[step] = strings.Replace(deps, doStep, "", -1)
	}
	return newDeps
}

func runPart1(input depType) {
	var result []string = []string{}
	var workingDeps depType
	workingDeps = input
	for len(workingDeps) > 0 {
		next := getNextStep(workingDeps)
		result = append(result, next)
		workingDeps = finishStep(workingDeps, next)
	}
	fmt.Println("part 1", strings.Join(result, ""))
}

type workerType struct {
	seconds int
	step    string
}

func workersDone(workers []workerType) bool {
	var done bool = true
	for i := range workers {
		if workers[i].step != "" {
			done = false
			break
		}
	}
	return done
}

func runPart2(input depType) {
	var workingDeps depType
	var workers []workerType = make([]workerType, 5)
	var t int = -1
	var done string
	var doing string

	workingDeps = input
	fmt.Printf("%6s %5s %5s %5s %5s %5s\n", "Time", "w1", "w2", "w3", "w4", "w5")
	for len(workingDeps) > 0 || !workersDone(workers) {
		t++
		for i := range workers {
			if workers[i].seconds == 0 {
				if workers[i].step != "" {
					workingDeps = finishStep(workingDeps, workers[i].step)
					done = done + workers[i].step
					doing = strings.Replace(doing, workers[i].step, "", -1)
					workers[i].step = ""
				}
			}
		}
		fmt.Printf("%6d", t)
		for i := range workers {
			if workers[i].seconds == 0 {
				if next := getNextStep(workingDeps); next != "" {
					workers[i].seconds = int(next[0]-'A') + 60
					workers[i].step = next
					workingDeps = startStep(workingDeps, next)
				}
			} else {
				workers[i].seconds--
			}
			fmt.Printf(" %1s(%2d)", workers[i].step, workers[i].seconds)
		}
		fmt.Print("\n")
	}
	fmt.Println("part 2 ", t)
}

func main() {
	input := getInput()
	runPart1(input)
	input = getInput()
	runPart2(input)
}
