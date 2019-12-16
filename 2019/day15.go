package main

import (
	"bufio"
	"fmt"
	"math"
	"time"
)

func init() {
	AddSolution(15, solveDay15)
}

func solveDay15(pr *PuzzleRun) {
	fmt.Print("\x1b[2J\x1b[H")
	scanner := bufio.NewScanner(pr.InFile)
	program := ""
	for scanner.Scan() {
		program = program + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	type move struct {
		dir int
		pos Point
	}

	comp, err := NewIntcodeFromInput(program)
	if err != nil {
		pr.logger.Fatal(err)
	}
	comp.AddStandardOpcodes()
	in := make(chan int64)
	out := make(chan int64)
	errchan := make(chan error)
	done := make(chan bool)

	comp.RunProgramChannelMode(in, out, errchan, done)

	pos := Point{0, 0}
	o2 := Point{}
	min := math.MaxInt64
	grid := NewGrid()
	grid.SetPoint(pos, '.')
	grid.SetBlank('`')
	path := []move{move{}}

	glyphs := []rune{'#', '.', 'O'}

	reverse := []int64{2, 1, 4, 3}
	for {
		dirs := []Point{
			pos.Add(North),
			pos.Add(South),
			pos.Add(West),
			pos.Add(East),
		}
		moved := false
		for i, d := range dirs {
			c := grid.GetPoint(d)
			if c != '`' { // seen that point before
				continue
			}
			in <- int64(i + 1)
			result := <-out
			grid.SetPoint(d, glyphs[result])
			if result == 0 { // hit a wall, try another direction
				continue
			}
			if result == 2 { // found the o² system
				o2 = d
				if min > len(path) {
					min = len(path)
				}
			}
			pos = d
			path = append(path, move{pos: pos, dir: i})
			moved = true
			break
		}
		if moved {
			continue
		} else if (pos == Point{0, 0}) {
			break
		}
		in <- reverse[path[len(path)-1].dir]
		<-out // throw away that output, it will be empty space
		pos = path[len(path)-2].pos
		path = path[:len(path)-1]
	}
	pr.ReportPart(min, o2)
	fmt.Print("\x1b[s")
	grid.SetPoint(Point{0, 0}, 'X')
	fmt.Print(grid)
	grid.SetPoint(Point{0, 0}, '.')
	//fmt.Print("\x1b[u")

	t := -1
	o2stack := []Point{o2}
	for len(o2stack) > 0 {
		newStack := []Point{}
		for _, point := range o2stack {
			grid.SetPoint(point, 'O')
		}
		for _, point := range o2stack {
			dirs := []Point{
				point.Add(North),
				point.Add(South),
				point.Add(West),
				point.Add(East),
			}
			for _, dir := range dirs {
				if grid.GetPoint(dir) == '.' {
					newStack = append(newStack, dir)
				}
			}
		}
		o2stack = newStack
		t++
		fmt.Print("\x1b[u")
		fmt.Print(grid)
		time.Sleep(10 * time.Millisecond)
	}
	pr.ReportPart(t)
}
