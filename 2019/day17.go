package main

import (
	"bufio"
	"fmt"
	"strings"
)

func init() {
	AddSolution(17, solveDay17)
}

func solveDay17(pr *PuzzleRun) {
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

	comp, err := NewIntcodeFromInput(program)
	pr.CheckError(err)
	comp.AddStandardOpcodes()
	outputs, err := comp.RunProgram(nil)
	grid := NewGrid()
	cursor := Point{0, 0}
	for _, n := range outputs {
		if n == 10 {
			cursor = Point{X: 0, Y: cursor.Y + 1}
			continue
		}
		grid.SetPoint(cursor, rune(n))
		cursor = Point{X: cursor.X + 1, Y: cursor.Y}
	}
	fmt.Println(grid)
	min, max := grid.Bounds()
	sum := 0
	dir := North
	bot := Point{0, 0}
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			point := Point{X: x, Y: y}
			c := grid.GetPoint(point)
			if c == '^' || c == 'v' || c == '<' || c == '>' {
				bot = point
				if c == '^' {
					dir = North
				}
				if c == 'v' {
					dir = South
				}
				if c == '<' {
					dir = West
				}
				if c == '>' {
					dir = East
				}
			}
			if c != '#' {
				continue
			}
			if grid.GetPoint(point.Add(North)) != '#' {
				continue
			}
			if grid.GetPoint(point.Add(South)) != '#' {
				continue
			}
			if grid.GetPoint(point.Add(East)) != '#' {
				continue
			}
			if grid.GetPoint(point.Add(West)) != '#' {
				continue
			}
			sum += point.X * point.Y
		}
	}
	pr.ReportPart(sum)

	count := 0
	sides := map[Point]map[rune]Point{
		North: map[rune]Point{'L': West, 'R': East},
		South: map[rune]Point{'L': East, 'R': West},
		East:  map[rune]Point{'L': North, 'R': South},
		West:  map[rune]Point{'L': South, 'R': North},
	}
	moves := make([]string, 0)
	pos := bot
	for {
		if grid.GetPoint(pos.Add(dir)) == '#' {
			pos = pos.Add(dir)
			count++
		} else {
			dirs := sides[dir]
			sidePoints := map[rune]Point{'L': pos.Add(dirs['L']), 'R': pos.Add(dirs['R'])}
			if count != 0 {
				moves = append(moves, fmt.Sprintf("%d", count))
			}
			count = 0
			if grid.GetPoint(sidePoints['L']) == '#' {
				moves = append(moves, "L")
				dir = dirs['L']
			} else if grid.GetPoint(sidePoints['R']) == '#' {
				moves = append(moves, "R")
				dir = dirs['R']
			} else {
				break
			}
		}
	}

	seq := strings.Join(moves, ",")
	fmt.Println(seq)

	// TODO: compress in code
	// sequence: L,10,L,8,R,8,L,8,R,6,L,10,L,8,R,8,L,8,R,6,R,6,R,8,R,8,R,6,R,6,L,8,L,10,R,6,R,8,R,8,R,6,R,6,L,8,L,10,R,6,R,8,R,8,R,6,R,6,L,8,L,10,R,6,R,8,R,8,L,10,L,8,R,8,L,8,R,6
	// Regex: (not a go regex, it uses backrefs) ^(.{3,20})(?:,\1)*,(.{3,20})(?:,(?:\1|\2))*,(.{3,20})(?:,(?:\1|\2|\3))*$
	a := "L,10,L,8,R,8,L,8,R,6"
	b := "R,6,R,8,R,8"
	c := "R,6,R,6,L,8,L,10"
	mainfunc := "A,A,B,C,B,C,B,C,B,A"

	comp, err = NewIntcodeFromInput(program)
	pr.CheckError(err)
	comp.AddStandardOpcodes()
	comp.mem[0] = 2

	in := make([]int64, 0, 5+len(mainfunc)+len(a)+len(b)+len(c))
	for _, r := range mainfunc {
		in = append(in, int64(r))
	}
	in = append(in, 10)
	for _, r := range a {
		in = append(in, int64(r))
	}
	in = append(in, 10)
	for _, r := range b {
		in = append(in, int64(r))
	}
	in = append(in, 10)
	for _, r := range c {
		in = append(in, int64(r))
	}
	in = append(in, 10)
	in = append(in, 'n')
	in = append(in, 10)
	fmt.Println(in)
	out, err := comp.RunProgram(in)
	pr.ReportPart(out[len(out)-1])
}
