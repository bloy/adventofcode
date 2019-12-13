package main

import (
	"bufio"
	"regexp"
	"strconv"
)

func init() {
	AddSolution(12, solveDay12)
}

// IntPoint is a point in N space
type IntPoint []int

type day12Moon struct {
	Position IntPoint
	Velocity IntPoint
}

// Pairs returns a list of tuples representing each possible pair of integers between 0 and i not inclusive of i
func Pairs(i int) [][]int {
	pairs := make([][]int, 0)
	for n := 0; n < i-1; n++ {
		for m := n + 1; m < i; m++ {
			pairs = append(pairs, []int{n, m})
		}
	}
	return pairs
}

// IntAbs returns the integer absolute value of i
func IntAbs(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}

func solveDay12(pr *PuzzleRun) {
	reg := regexp.MustCompile(`^<x=([-0-9]+), y=([-0-9]+), z=([-0-9]+)>$`)
	moons := make([]*day12Moon, 0)
	scanner := bufio.NewScanner(pr.InFile)
	for scanner.Scan() {
		strs := reg.FindStringSubmatch(scanner.Text())[1:]
		nums := make([]int, len(strs))
		for i := range strs {
			num, err := strconv.Atoi(strs[i])
			if err != nil {
				pr.logger.Fatal(err)
			}
			nums[i] = num
		}
		moon := day12Moon{}
		moon.Position = nums
		moon.Velocity = IntPoint{0, 0, 0}
		moons = append(moons, &moon)
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	pairs := Pairs(len(moons))
	for i := 0; i < 1000; i++ {
		for _, pair := range pairs {
			m0 := moons[pair[0]]
			m1 := moons[pair[1]]
			for j := 0; j < 3; j++ {
				if m0.Position[j] > m1.Position[j] {
					m0.Velocity[j]--
					m1.Velocity[j]++
				} else if m0.Position[j] < m1.Position[j] {
					m0.Velocity[j]++
					m1.Velocity[j]--
				}
			}
		}
		for _, moon := range moons {
			for j := 0; j < 3; j++ {
				moon.Position[j] += moon.Velocity[j]
			}
		}
	}

	energy := 0
	for _, moon := range moons {
		potential := 0
		kenetic := 0
		for j := 0; j < 3; j++ {
			potential += IntAbs(moon.Position[j])
			kenetic += IntAbs(moon.Velocity[j])
		}
		energy += potential * kenetic
	}
	pr.ReportPart("total kenetic energy after 1000 steps:", energy)
}
