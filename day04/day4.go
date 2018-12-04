package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func getInput() []string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(content), "\n")
	sort.Strings(strs)
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
