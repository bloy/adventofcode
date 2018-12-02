package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getInput() []string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(content), "\n")
	return strs
}

func runPart1(input []string) {
	count2 := 0
	count3 := 0
	for _, str := range input {
		has2 := false
		has3 := false
		for i := range str {
			count := strings.Count(str, str[i:i+1])
			switch count {
			case 2:
				has2 = true
			case 3:
				has3 = true
			}
		}
		if has2 {
			count2 += 1
		}
		if has3 {
			count3 += 1
		}
	}
	answer := count2 * count3
	fmt.Println("part 1 ", answer)
}

func runPart2(input []string) {
	var finalStr1, finalStr2 string
	var differences int
	for _, str1 := range input {
		if str1 == "" {
			continue
		}
		for _, str2 := range input {
			if str2 == "" {
				continue
			}
			differences = 0
			for i := range str1 {
				if str1[i] != str2[i] {
					differences += 1
				}
			}
			if differences == 1 {
				fmt.Printf("Found strings %s and %s\n", str1, str2)
				finalStr1 = str1
				finalStr2 = str2
			}
		}
	}
	fmt.Printf("part 2 answer: ")
	for i := range finalStr1 {
		if finalStr1[i] == finalStr2[i] {
			fmt.Print(finalStr1[i : i+1])
		}
	}
	fmt.Println("")
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
