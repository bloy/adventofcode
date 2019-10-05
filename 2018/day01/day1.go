package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput() []int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(content), "\n")
	ints := make([]int, 0, len(strs))
	for _, str := range strs {
		if str == "" {
			continue
		}
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func runPart1(input []int) {
	answer := 0
	for _, i := range input {
		answer += i
	}
	fmt.Println("part 1 ", answer)
}

func runPart2(input []int) {
	seen := make(map[int]int)
	freq := 0
	seen[0] = 1
	for {
		for _, i := range input {
			freq += i
			seen[freq] = seen[freq] + 1
			if seen[freq] == 2 {
				answer := freq
				fmt.Println("part 2 ", answer)
				return
			}
		}
	}
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
