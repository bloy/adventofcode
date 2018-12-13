package main

import (
	"container/list"
	"fmt"
	"strings"
)

const inputState = `.#####.##.#.##...#.#.###..#.#..#..#.....#..####.#.##.#######..#...##.#..#.#######...#.#.#..##..#.#.#`
const rulesStrs = `#..#. => .
##... => #
#.... => .
#...# => #
...#. => .
.#..# => #
#.#.# => .
..... => .
##.## => #
##.#. => #
###.. => #
#.##. => .
#.#.. => #
##..# => #
..#.# => #
..#.. => .
.##.. => .
...## => #
....# => .
#.### => #
#..## => #
..### => #
####. => #
.#.#. => #
.#### => .
###.# => #
##### => #
.#.## => .
.##.# => .
.###. => .
..##. => .
.#... => #`

func getInput() (initial *list.List, rules map[string]bool) {
	initial = list.New()
	rules = make(map[string]bool)
	for _, c := range inputState {
		switch c {
		case '.':
			initial.PushBack(false)
		case '#':
			initial.PushBack(true)
		}
	}
	for _, str := range strings.Split(rulesStrs, "\n") {
		parts := strings.Split(str, " ")
		if parts[2] == "#" {
			rules[parts[0]] = true
		} else {
			rules[parts[0]] = false
		}
	}
	return
}

func printPots(firstPot int, state *list.List) {
	fmt.Printf("%d ", firstPot)
	for pot := state.Front(); pot != nil; pot = pot.Next() {
		if pot.Value.(bool) {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("")
}

func runPart1(state *list.List, rules map[string]bool) {
	firstPot := 0
	pots := list.New()
	pots.PushBackList(state)
	printPots(firstPot, pots)
	for generation := 0; generation < 20; generation++ {
		nextPots := list.New()
		for extra := 0; extra < 4; extra++ {
			pots.PushFront(false)
			pots.PushBack(false)
		}
		potNum := firstPot - 2
		filledPot := false
		for pot := pots.Front().Next().Next(); pot.Next().Next() != nil; pot = pot.Next() {
			checks := make([]string, 0, 5)
			rangePot := pot.Prev().Prev()
			for i := 0; i < 5; i++ {
				var checkchar string
				if rangePot.Value.(bool) {
					checkchar = "#"
				} else {
					checkchar = "."
				}
				checks = append(checks, checkchar)
				rangePot = rangePot.Next()
			}
			checkStr := strings.Join(checks, "")
			if rules[checkStr] == true {
				nextPots.PushBack(true)
				if !filledPot {
					filledPot = true
					firstPot = potNum
				}
			} else if filledPot { // previous filled pot and this one is not filled
				nextPots.PushBack(false)
			}
			potNum++
		}
		pots = nextPots
		printPots(firstPot, pots)
	}
	sum := 0
	num := firstPot
	for e := pots.Front(); e != nil; e = e.Next() {
		if e.Value.(bool) {
			sum += num
		}
		num++
	}
	fmt.Println("part 1", firstPot, sum)
}

func runPart2(state *list.List, rules map[string]bool) {
	firstPot := 0
	pots := list.New()
	pots.PushBackList(state)
	printPots(firstPot, pots)
	for generation := 0; generation < 50000000000; generation++ {
		if generation%100 == 0 {
			var genf float32 = float32(generation)
			var pct float32 = genf / 50000000000 * 100
			fmt.Printf("gen %d - %0.2f percent\n", generation, pct)
		}
		nextPots := list.New()
		for extra := 0; extra < 4; extra++ {
			pots.PushFront(false)
			pots.PushBack(false)
		}
		potNum := firstPot - 2
		filledPot := false
		for pot := pots.Front().Next().Next(); pot.Next().Next() != nil; pot = pot.Next() {
			checks := make([]string, 0, 5)
			rangePot := pot.Prev().Prev()
			for i := 0; i < 5; i++ {
				var checkchar string
				if rangePot.Value.(bool) {
					checkchar = "#"
				} else {
					checkchar = "."
				}
				checks = append(checks, checkchar)
				rangePot = rangePot.Next()
			}
			checkStr := strings.Join(checks, "")
			if rules[checkStr] == true {
				nextPots.PushBack(true)
				if !filledPot {
					filledPot = true
					firstPot = potNum
				}
			} else if filledPot { // previous filled pot and this one is not filled
				nextPots.PushBack(false)
			}
			potNum++
		}
		pots = nextPots
	}
	sum := 0
	num := firstPot
	for e := pots.Front(); e != nil; e = e.Next() {
		if e.Value.(bool) {
			sum += num
		}
		num++
	}
	fmt.Println("part 2", firstPot, sum)
}

func main() {
	state, rules := getInput()
	runPart1(state, rules)
	//runPart2(state, rules)
}
