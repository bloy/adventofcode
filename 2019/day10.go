package main

import (
	"bufio"
	"math"
	"sort"
	"time"

	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
)

func init() {
	AddSolution(10, solveDay10)
}

func solveDay10(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	grid := NewGrid()
	asteroids := []Point{}
	y := -1
	for scanner.Scan() {
		y++
		line := scanner.Text()
		for x, r := range line {
			if r == '#' {
				p := Point{X: x, Y: y}
				grid.SetPoint(p, r)
				asteroids = append(asteroids, p)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	var (
		bestNum   int
		bestPoint Point
		count     int
	)

	for _, base := range asteroids {
		count = 0
		seen := make(map[float64]bool)
		for _, a := range asteroids {
			if a == base {
				continue
			}
			d := Point{X: a.X - base.X, Y: a.Y - base.Y}
			at := math.Atan2(float64(d.Y), float64(d.X))
			if seen[at] {
				continue
			}
			seen[at] = true
			count++
		}
		if count > bestNum {
			bestNum = count
			bestPoint = base
		}
	}
	pr.ReportPart("Point", bestPoint, "Visible", bestNum, "Total", len(asteroids))

	base := bestPoint
	angles := make([]float64, 0)
	byAngle := make(map[float64][]Point)
	for _, a := range asteroids {
		if a == base {
			continue
		}
		d := Point{X: a.X - base.X, Y: a.Y - base.Y}
		at := math.Atan2(float64(d.Y), float64(d.X)) + math.Pi/2
		if at < 0 {
			at = at + math.Pi*2
		}
		if _, ok := byAngle[at]; !ok {
			angles = append(angles, at)
			byAngle[at] = make([]Point, 0)
		}
		byAngle[at] = append(byAngle[at], a)
	}
	sort.Slice(angles, func(i, j int) bool {
		return angles[i] < angles[j]
	})
	for a := range byAngle {
		sort.Slice(byAngle[a], func(i, j int) bool {
			pi := byAngle[a][i]
			pj := byAngle[a][j]
			di := Point{X: pi.X - base.X, Y: pi.Y - base.Y}
			dj := Point{X: pj.X - base.X, Y: pj.Y - base.Y}
			return math.Hypot(float64(di.X), float64(di.Y)) < math.Hypot(float64(dj.X), float64(dj.Y))
		})
	}

	destroyed := make([]Point, 0, len(asteroids))
	_, maxPoint := grid.Bounds()
	grid.SetPoint(base, '@')
	color.Output = ansi.NewAnsiStdout()
	grid.AddRuneColor('.', color.New(color.FgHiBlack))
	grid.AddRuneColor('@', color.New(color.FgHiYellow))
	grid.AddRuneColor('#', color.New(color.FgRed))
	grid.ColorPrint()
	time.Sleep(time.Millisecond * 500)

	var target Point
	for len(destroyed) < len(asteroids)-1 {
		for _, angle := range angles {
			if len(byAngle[angle]) > 0 {
				a := byAngle[angle][0]
				byAngle[angle] = byAngle[angle][1:]
				destroyed = append(destroyed, a)
				grid.SetPoint(a, '.')
				ansi.CursorPreviousLine(maxPoint.Y + 1)
				grid.ColorPrint()
				if len(destroyed) == 200 {
					target = a
				}
				time.Sleep(time.Millisecond * 10)
			}
		}
	}
	pr.ReportPart("200th:", target, target.X*100+target.Y)
}
