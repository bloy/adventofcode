package main

import (
	"bufio"
	"strings"
)

func init() {
	AddSolution(20, solveDay20)
}
func solveDay20(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	b := &strings.Builder{}
	for scanner.Scan() {
		b.WriteString(scanner.Text())
		b.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		pr.CheckError(err)
	}
	mapgrid := b.String()

	grid := day20buildMaze(mapgrid)
	pr.ReportLoad()

	pr.logger.Println(grid.Labels)
	pr.ReportPart(day20solveMaze(grid))
}

type day20grid struct {
	Grid
	Portals map[Point]Point
	Labels  map[string][]Point
}

func day20solveMaze(grid *day20grid) int {
	dist := make(map[Point]int)
	seen := make(map[Point]bool)
	q := make([]Point, 1)
	goal := grid.Labels["ZZ"][0]
	start := grid.Labels["AA"][0]
	q[0] = start
	dist[start] = 0
	seen[start] = true
	found := false
	for len(q) > 0 && !found {
		current := q[0]
		q = q[1:]
		adjPoints := []Point{
			current.Add(North),
			current.Add(East),
			current.Add(West),
			current.Add(South),
		}
		if p, ok := grid.Portals[current]; ok {
			adjPoints = append(adjPoints, p)
		}
		for _, adj := range adjPoints {
			if seen[adj] {
				continue
			}
			c := grid.GetPoint(adj)
			if c == '#' || c == ' ' {
				continue
			}
			seen[adj] = true
			dist[adj] = dist[current] + 1
			q = append(q, adj)
			if adj == goal {
				found = true
				break
			}
		}
	}
	if !found {
		return 0
	}
	return dist[goal]
}

func day20gridRecordLabel(g *day20grid, p Point, labelseen map[Point]bool) {
	if labelseen[p] {
		return
	}
	c1 := g.GetPoint(p)
	var p2 Point
	var spot Point
	if g.GetPoint(p.Add(North)) == '.' {
		spot = p.Add(North)
		p2 = p.Add(South)
	} else if g.GetPoint(p.Add(West)) == '.' {
		spot = p.Add(West)
		p2 = p.Add(East)
	} else if g.GetPoint(p.Add(South).Add(South)) == '.' {
		spot = p.Add(South).Add(South)
		p2 = p.Add(South)
	} else {
		spot = p.Add(East).Add(East)
		p2 = p.Add(East)
	}
	labelseen[p] = true
	labelseen[p2] = true
	c2 := g.GetPoint(p2)
	label := string(c1) + string(c2)
	g.Labels[label] = append(g.Labels[label], spot)
}

func day20buildMaze(mapgrid string) *day20grid {
	g := &day20grid{}
	g.Portals = make(map[Point]Point)
	g.Labels = make(map[string][]Point)
	g.Grid = *NewGridFromInput(mapgrid)
	g.SetBlank(' ')
	labelseen := make(map[Point]bool)
	min, max := g.Bounds()
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			p := Point{X: x, Y: y}
			c := g.GetPoint(p)
			if c == ' ' {
				continue
			}
			if c >= 'A' && c <= 'Z' {
				day20gridRecordLabel(g, p, labelseen)
			}
		}
	}
	for _, list := range g.Labels {
		if len(list) == 2 {
			g.Portals[list[0]] = list[1]
			g.Portals[list[1]] = list[0]
		}
	}
	return g
}

const (
	day20test1 = `         A           
         A           
  #######.#########  
  #######.........#  
  #######.#######.#  
  #######.#######.#  
  #######.#######.#  
  #####  B    ###.#  
BC...##  C    ###.#  
  ##.##       ###.#  
  ##...DE  F  ###.#  
  #####    G  ###.#  
  #########.#####.#  
DE..#######...###.#  
  #.#########.###.#  
FG..#########.....#  
  ###########.#####  
             Z       
             Z       `

	day20test2 = `                   A               
                   A               
  #################.#############  
  #.#...#...................#.#.#  
  #.#.#.###.###.###.#########.#.#  
  #.#.#.......#...#.....#.#.#...#  
  #.#########.###.#####.#.#.###.#  
  #.............#.#.....#.......#  
  ###.###########.###.#####.#.#.#  
  #.....#        A   C    #.#.#.#  
  #######        S   P    #####.#  
  #.#...#                 #......VT
  #.#.#.#                 #.#####  
  #...#.#               YN....#.#  
  #.###.#                 #####.#  
DI....#.#                 #.....#  
  #####.#                 #.###.#  
ZZ......#               QG....#..AS
  ###.###                 #######  
JO..#.#.#                 #.....#  
  #.#.#.#                 ###.#.#  
  #...#..DI             BU....#..LF
  #####.#                 #.#####  
YN......#               VT..#....QG
  #.###.#                 #.###.#  
  #.#...#                 #.....#  
  ###.###    J L     J    #.#.###  
  #.....#    O F     P    #.#...#  
  #.###.#####.#.#####.#####.###.#  
  #...#.#.#...#.....#.....#.#...#  
  #.#####.###.###.#.#.#########.#  
  #...#.#.....#...#.#.#.#.....#.#  
  #.###.#####.###.###.#.#.#######  
  #.#.........#...#.............#  
  #########.###.###.#############  
           B   J   C               
           U   P   P               `
)
