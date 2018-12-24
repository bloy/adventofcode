package main

import "fmt"

func main() {
	allGroups := parseInput(inputStr)
	for _, g := range allGroups {
		fmt.Println(g)
	}
}
