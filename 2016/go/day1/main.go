package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type dataType []instruction

type point struct {
	x, y int
}

var (
	north = point{x: 0, y: -1}
	south = point{x: 0, y: 1}
	east  = point{x: 1, y: 0}
	west  = point{x: -1, y: 0}
)

var facingMachine = map[point]map[string]point{
	north: map[string]point{"L": west, "R": east},
	south: map[string]point{"L": east, "R": west},
	east:  map[string]point{"L": north, "R": south},
	west:  map[string]point{"L": south, "R": north},
}

type instruction struct {
	turn string
	num  int
}

func (i instruction) doTurn(dir point) point {
	newDir := facingMachine[dir][i.turn]
	return newDir
}

func (i instruction) doStep(dir point, cur point) (newdir point, newpoint point) {
	newdir = i.doTurn(dir)
	newpoint.x = cur.x + (newdir.x * i.num)
	newpoint.y = cur.y + (newdir.y * i.num)
	return
}

func (p point) absValue() int {
	v := 0
	if p.x < 0 {
		v += (p.x * -1)
	} else {
		v += p.x
	}
	if p.y < 0 {
		v += (p.y * -1)
	} else {
		v += p.y
	}
	return v
}

func main() {
	fmt.Println("Advent of code 2016 Day 1")
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

func readInput(inputText string) dataType {
	strs := strings.Split(inputText, ", ")
	instrs := make(dataType, len(strs))
	for i, s := range strs {
		d, n := s[:1], s[1:]
		m, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		instrs[i] = instruction{turn: d, num: m}
	}
	return instrs
}

func solve1(data dataType) int {
	dir := north
	cur := point{0, 0}
	for _, instr := range data {
		dir, cur = instr.doStep(dir, cur)
	}
	return cur.absValue()
}

func solve2(data dataType) int {
	dir := north
	cur := point{0, 0}
	seen := make(map[point]bool)
	seen[cur] = true
	for _, instr := range data {
		dir = instr.doTurn(dir)
		for i := 0; i < instr.num; i++ {
			cur.x = cur.x + dir.x
			cur.y = cur.y + dir.y
			if seen[cur] {
				return cur.absValue()
			}
			seen[cur] = true
		}
	}
	return cur.absValue()
}
