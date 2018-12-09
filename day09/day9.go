package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const testStr string = ``

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func getInput() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(content)
	str = testStr
	return str
}

func runPart1(input string) {
	fmt.Println("part 1", input)
}

func runPart2(input string) {
	fmt.Println("part 2")
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
