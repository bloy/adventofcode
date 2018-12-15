package main

import "fmt"
import "strconv"

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func runPart1(numRecipes int) {
	var recipes []byte = make([]byte, 2)
	var elves [2]int = [2]int{0, 1}
	recipes[0] = 3
	recipes[1] = 7
	for len(recipes) < numRecipes+10 {
		combine := recipes[elves[0]] + recipes[elves[1]]
		if combine/10 != 0 {
			recipes = append(recipes, combine/10)
		}
		recipes = append(recipes, combine%10)
		for i := range elves {
			elves[i] = (elves[i] + 1 + int(recipes[elves[i]])) % len(recipes)
		}
	}
	fmt.Printf("Input: %6d  Output: ", numRecipes)
	for i := numRecipes; i < numRecipes+10; i++ {
		fmt.Printf("%d", recipes[i])
	}
	fmt.Print("\n")
}

func compareSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		panic("oops")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func runPart2(pattern string) {
	var recipes []int = make([]int, 2)
	var elves [2]int = [2]int{0, 1}
	var count int = 2
	var expected []int = make([]int, 0, len(pattern))
	recipes[0] = 3
	recipes[1] = 7
	for i := 0; i < len(pattern); i++ {
		expected = append(expected, atoi(string(pattern[i])))
	}
	for len(recipes) < len(pattern)+2 || !compareSlices(expected, recipes[count-len(pattern):count]) {
		count++
		combine := recipes[elves[0]] + recipes[elves[1]]
		if combine/10 != 0 {
			recipes = append(recipes, combine/10)
		}
		recipes = append(recipes, combine%10)
		for i := range elves {
			elves[i] = (elves[i] + 1 + int(recipes[elves[i]])) % len(recipes)
		}
	}
	fmt.Printf("Input: %6s  Output: %d\n", pattern, count-len(pattern))
}

func main() {
	runPart1(9)
	runPart1(5)
	runPart1(18)
	runPart1(2018)
	runPart1(554401)
	runPart2("51589")
	runPart2("01245")
	runPart2("92510")
	runPart2("59414")
	runPart2("554401")
}
