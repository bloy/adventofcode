package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getInput() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(content)
	str = `dabAcCaCBAcCcaDA`
	return str
}

func runPart1(input string) {
	fmt.Println("part 1 ")
}

func runPart2(input string) {
	fmt.Println("part 2 ")
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
