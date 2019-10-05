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

var rules map[string]bool

func getInput() (initial *list.List) {
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

func printPots(firstPot int, pots *list.List) {
	fmt.Printf("%d ", firstPot)
	for pot := pots.Front(); pot != nil; pot = pot.Next() {
		if pot.Value.(bool) {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("")
}

func sumPots(firstPot int, pots *list.List) int {
	sum := 0
	num := firstPot
	for e := pots.Front(); e != nil; e = e.Next() {
		if e.Value.(bool) {
			sum += num
		}
		num++
	}
	return sum
}

func iteratePots(firstPot int, pots *list.List) (newFirstPot int, nextPots *list.List) {
	for extra := 0; extra < 4; extra++ {
		pots.PushFront(false)
		pots.PushBack(false)
	}
	newFirstPot = firstPot - 2
	nextPots = list.New()
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
		nextPots.PushBack(rules[checkStr])
	}
	for nextPots.Front().Value.(bool) == false {
		nextPots.Remove(nextPots.Front())
		newFirstPot++
	}
	for nextPots.Back().Value.(bool) == false {
		nextPots.Remove(nextPots.Back())
	}
	return
}

func runPart1(state *list.List) {
	firstPot := 0
	pots := list.New()
	pots.PushBackList(state)
	printPots(firstPot, pots)
	for generation := 0; generation < 20; generation++ {
		firstPot, pots = iteratePots(firstPot, pots)
		printPots(firstPot, pots)
	}
	fmt.Println("part 1", firstPot, sumPots(firstPot, pots))
}

func runPart2(state *list.List) {
	limit := 50000000000
	firstPot := 0
	pots := list.New()
	pots.PushBackList(state)
	printPots(firstPot, pots)
	prevSum := 0
	prevDiff := 0
	generation := 0
	diff := -1
	for prevDiff != diff || generation < 100 { // 100 is enough to establish the pattern
		generation++
		prevDiff = diff
		firstPot, pots = iteratePots(firstPot, pots)
		sum := sumPots(firstPot, pots)
		diff = sum - prevSum
		prevSum = sum
		fmt.Printf("Generation %d, firstPot %d, sum %d (%d)\n", generation, firstPot, sum, diff)
	}
	total := (limit-generation)*prevDiff + prevSum
	fmt.Println("part 2", total)
}

func main() {
	state := getInput()
	runPart1(state)
	runPart2(state)
}
