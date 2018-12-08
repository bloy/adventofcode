package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type depType map[string]string

func getInput() depType {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(strings.TrimSpace(string(content)), "\n")
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

func runPart2(input depType) {
	//var result []string = []string{}
	//var workingDeps depType
	//var workers [5]workerType
	//var t int

	//workingDeps = input

	fmt.Println("part 2")
}

func main() {
	input := getInput()
	runPart1(input)
	input = getInput()
	runPart2(input)
}
