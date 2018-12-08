package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const testStr string = `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`

func getInput() []int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(content)
	str = testStr
	strs := strings.Split(strings.TrimSpace(str), " ")
	var intList []int = []int{}
	for _, str := range strs {
		if i, err := strconv.Atoi(str); err == nil {
			intList = append(intList, i)
		} else {
			panic(err)
		}
	}
	return intList
}

func runPart1(input []int) {
	fmt.Println("part 1", " n/a")
}

func runPart2(input []int) {
	fmt.Println("part 2", " n/a")
}

func main() {
	input := getInput()
	runPart1(input)
	input = getInput()
	runPart2(input)
}
