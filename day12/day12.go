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
		rules[parts[0]] = rules[parts[2]]
	}
	return
}

func runPart1(state *list.List, rules map[string]bool) {
	firstPot := 0
	for i := 0; i < 20; i++ {
		nextState := list.New()
		minus1pots := []string{".", ".", ".", ".", "."}
		minus2pots := []string{".", ".", ".", ".", "."}
		if state.Front() != nil {
			if state.Front().Value.(bool) {
				minus1pots[3] = "#"
				minus2pots[4] = "#"
			}
			if state.Front().Next() != nil && state.Front().Next().Value.(bool) {
				minus1pots[4] = "#"
			}
		}
		minus1check := strings.Join(minus1pots, "")
		minus2check := strings.Join(minus2pots, "")
		if rules[minus1check] == true && rules[minus2check] == true {
			nextState.PushFront(true)
			nextState.PushFront(true)
			firstPot -= 2
		} else if rules[minus1check] == true {
			nextState.PushFront(true)
			firstPot -= 1
		} else if rules[minus2check] == true {
			nextState.PushFront(false)
			nextState.PushFront(true)
			firstPot -= 2
		}
		for e := state.Front(); e != nil; e = e.Next() {
			pots := []string{".", ".", ".", ".", "."}
			if e.Prev() != nil {
				if e.Prev().Value.(bool) {
					pots[1] = "#"
				}
				if e.Prev().Prev() != nil && e.Prev().Prev().Value.(bool) {
					pots[0] = "#"
				}
			}
			if e.Value.(bool) {
				pots[2] = "#"
			}
			if e.Next() != nil {
				if e.Next().Value.(bool) {
					pots[3] = "#"
				}
				if e.Next().Next() != nil && e.Next().Next().Value.(bool) {
					pots[4] = "#"
				}
			}
			check := strings.Join(pots, "")
			if rules[check] == true {
				nextState.PushBack(true)
			} else {
				nextState.PushBack(false)
			}
		}
		plus1pots := []string{".", ".", ".", ".", "."}
		plus2pots := []string{".", ".", ".", ".", "."}
		if state.Back() != nil {
			if state.Back().Value.(bool) {
				plus1pots[1] = "#"
				plus2pots[0] = "#"
			}
			if state.Back().Prev() != nil && state.Back().Prev().Value.(bool) {
				plus1pots[0] = "#"
			}
		}
		plus1check := strings.Join(plus1pots, "")
		plus2check := strings.Join(plus2pots, "")
		if rules[plus1check] == true && rules[plus2check] == true {
			nextState.PushBack(true)
			nextState.PushBack(true)
		} else if rules[plus1check] == true {
			nextState.PushBack(true)
		} else if rules[plus2check] == true {
			nextState.PushBack(false)
			nextState.PushBack(true)
		}
		state = nextState
	}
	sum := 0
	num := firstPot
	for e := state.Front(); e != nil; e = e.Next() {
		if e.Value.(bool) {
			sum += num
		}
		num++
	}
	fmt.Println("part 1", firstPot, sum)
}

func runPart2(state *list.List, rules map[string]bool) {
	fmt.Println("part 2")
}

func main() {
	state, rules := getInput()
	runPart1(state, rules)
	runPart2(state, rules)
}
