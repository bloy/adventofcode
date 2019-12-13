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

type day12axisstate struct {
	p1, p2, p3, p4 int
	v1, v2, v3, v4 int
}

// map between an axis state and the first time it was seen
type day12axisseen map[day12axisstate]int

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

func day12GCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a == b {
		return a
	}
	if a > b {
		return day12GCD(a-b, b)
	}
	return day12GCD(a, b-a)
}

func day12LCM(a, b int) int {
	return a / day12GCD(a, b) * b
}

func solveDay12(pr *PuzzleRun) {
	reg := regexp.MustCompile(`^<x=([-0-9]+), y=([-0-9]+), z=([-0-9]+)>$`)
	moons := make([]*day12Moon, 0)
	moons2 := make([]*day12Moon, 0)
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
		moon2 := day12Moon{}
		moon.Position = nums
		moon2.Position = nums
		moon.Velocity = IntPoint{0, 0, 0}
		moon2.Velocity = IntPoint{0, 0, 0}
		moons = append(moons, &moon)
		moons2 = append(moons, &moon2)
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

	moons = moons2
	step := 0
	looped := make([]bool, 3)
	seen := make([]day12axisseen, 3)
	loops := make([]int, 3)
	for i := 0; i < 3; i++ {
		seen[i] = make(day12axisseen)
	}
	for !looped[0] || !looped[1] || !looped[2] {
		for i := 0; i < 3; i++ {
			if !looped[i] {
				state := day12axisstate{
					p1: moons[0].Position[i], v1: moons[0].Velocity[i],
					p2: moons[1].Position[i], v2: moons[1].Velocity[i],
					p3: moons[2].Position[i], v3: moons[2].Velocity[i],
					p4: moons[2].Position[i], v4: moons[3].Velocity[i],
				}
				_, ok := seen[i][state]
				if ok {
					looped[i] = true
					loops[i] = step
					pr.logger.Printf("    %d loop 0 to %d", i, step)
				} else {
					seen[i][state] = step
				}
			}
		}
		step++
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

	L := day12LCM(day12LCM(loops[0], loops[1]), loops[2])
	pr.ReportPart(L)
}
