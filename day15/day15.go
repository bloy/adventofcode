package main

import "fmt"

func part1(levelStr string, debug bool) {
	var level *Level = NewLevel(levelStr)
	if debug {
		fmt.Println(level)
	}
}

func main() {
	part1(test1, true)
}
