package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const testStr string = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

func getInput() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(content)
	str = testStr
	str = strings.TrimSpace(str)
	return str
}

func runPart1(input string) {
	fmt.Println("part 1")
}

func runPart2(input string) {
	fmt.Println("part 2")
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
