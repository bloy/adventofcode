package main

import "fmt"

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
		//for i, val := range recipes {
		//c1 := " "
		//c2 := " "
		//if elves[0] == i && elves[1] == i {
		//c1 = "{"
		//c2 = "}"
		//} else if elves[1] == i {
		//c1 = "["
		//c2 = "]"
		//} else if elves[0] == i {
		//c1 = "("
		//c2 = ")"
		//}
		//fmt.Printf("%s%d%s", c1, val, c2)
		//}
		//fmt.Print("\n")
	}
	fmt.Printf("Input: %6d  Output: ", numRecipes)
	for i := numRecipes; i < numRecipes+10; i++ {
		fmt.Printf("%d", recipes[i])
	}
	fmt.Print("\n")
}

func main() {
	runPart1(9)
	runPart1(5)
	runPart1(18)
	runPart1(2018)
	runPart1(554401)
}
