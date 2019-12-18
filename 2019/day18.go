package main

import (
	"bufio"
	"strings"
)

func init() {
	AddSolution(18, solveDay18)
}

func solveDay18(pr *PuzzleRun) {
	lines := []string{}
	scanner := bufio.NewScanner(pr.InFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		pr.CheckError(err)
	}
	grid := NewFixedGrid(strings.Join(lines, "\n"))
	var pos Point
	keyBits := make(map[rune]int)
	doorBits := make(map[rune]int)
	var allKeys int
	for y := 0; y < grid.Size.Y; y++ {
		for x := 0; x < grid.Size.X; x++ {
			p := Point{X: x, Y: y}
			c := grid.GetPoint(p)
			if c == '@' {
				pos = p
			} else if c >= 'a' && c <= 'z' {
				keyBits[c] = 1 << (c - 'a')
				allKeys |= keyBits[c]
			} else if c >= 'A' && c <= 'Z' {
				doorBits[c] = 1 << (c - 'A')
			}
		}
	}
	pr.ReportLoad()
	pr.logger.Print(grid)
	pr.logger.Print(pos)
	pr.logger.Print(allKeys)
	pr.logger.Print(keyBits)
}
