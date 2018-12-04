package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Claim struct {
	num, x, y, width, height int
}

func getInput() []Claim {
	re := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(content), "\n")
	claims := make([]Claim, 0, len(strs))
	for _, str := range strs {
		if str == "" {
			continue
		}
		matches := re.FindStringSubmatch(str)
		data := make([]int, 0, len(matches)-1)
		for _, s := range matches[1:] {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			data = append(data, i)
		}
		claim := Claim{data[0], data[1], data[2], data[3], data[4]}
		claims = append(claims, claim)
	}
	return claims
}

func runPart1(input []Claim) {
	fabric := make(map[Point]int)
	for _, claim := range input {
		for x := claim.x; x < claim.x+claim.width; x++ {
			for y := claim.y; y < claim.y+claim.height; y++ {
				fabric[Point{x, y}] = fabric[Point{x, y}] + 1
			}
		}
	}
	counter := 0
	for _, value := range fabric {
		if value > 1 {
			counter++
		}
	}
	fmt.Println("part 1 ", counter)
}

func runPart2(input []Claim) {
	fmt.Println("part 1 ", "N/A")
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
