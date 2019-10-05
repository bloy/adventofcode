package main

import "fmt"

func doPart1(area *Area) {
	fmt.Println(area)
	for i := 0; i < 10; i++ {
		area.Step()
	}
	fmt.Println("Part 1: total value =", area.Value())
}

func doPart2(area *Area) {
	num := 1000000000
	loopstart := 0
	loopend := 0
	archive := make([]TileList, 0, 500)
	repeat := false
	for i := 0; i < num; i++ {
		area.Step()
		for old := 0; old < len(archive); old++ {
			if isSameArea(&(archive[old]), &(area.Tiles)) {
				fmt.Println("step", i, "repeats step", old)
				loopstart = old
				loopend = i
				repeat = true
				break
			}
		}
		if repeat {
			break
		}
		archive = append(archive, area.Tiles)
	}
	looplength := loopend - loopstart
	fmt.Println("loop start:", loopstart, "loop end:", loopend, "loop length:", looplength)
	stepNum := (((num - loopend) / looplength) * looplength) + loopend + 1
	area.Tiles = archive[loopstart]
	for i := stepNum; i < num; i++ {
		area.Step()
	}

	fmt.Println("Part 2: total value =", area.Value(), "after", num, "steps")
}

func isSameArea(aptr, bptr *TileList) bool {
	a := *aptr
	b := *bptr
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	area := parseInput(inputStr)
	doPart1(area)
	area = parseInput(inputStr)
	doPart2(area)
}
