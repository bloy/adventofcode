package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Advent of code 2016 Day 2")
	now := time.Now()
	start := now
	data := readInput(inputText)
	fmt.Println("  Data read:", time.Since(now))
	now = time.Now()
	fmt.Printf("  Part 1: %v (%s)\n", solve1(data), time.Since(now))
	now = time.Now()
	fmt.Printf("  Part 2: %v (%s)\n", solve2(data), time.Since(now))
	fmt.Println("  Total Time:", time.Since(start))
}
