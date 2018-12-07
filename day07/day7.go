package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func getInput() []string {
	re := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(strings.TrimSpace(string(content)), "\n")
	return strs
}

func runPart1(input []string) {
	fmt.Println("part 1 ")
}

func runPart2(input []string) {
	fmt.Println("part 2 ")
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
