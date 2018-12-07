package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type depType map[string][]string

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
		if deps[after] == nil {
			deps[after] = make([]string, 0, 1)
		}
		if deps[before] == nil {
			deps[before] = make([]string, 0, 1)
		}
		deps[before] = append(deps[before], after)
	}
	return deps
}

func runPart1(input depType) {
	fmt.Println(input)
	fmt.Println("part 1 ")
}

func runPart2(input depType) {
	fmt.Println("part 2 ")
}

func main() {
	input := getInput()
	runPart1(input)
	input = getInput()
	runPart2(input)
}
