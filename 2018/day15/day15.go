package main

import "fmt"

func part1(levelStr string, debug bool) {
	var level *Level = NewLevel(levelStr, 3)
	fmt.Println(level)
	for fullRound, _ := level.PlayRound(); fullRound; {
		if debug {
			fmt.Println(level)
		}
		fullRound, _ = level.PlayRound()
	}
	fmt.Println(level)
	fmt.Println(level.Score())
}

func part2(levelStr string, debug bool) {
	power := 3
	elfDead := true
	for elfDead {
		power++
		fmt.Println("Simulating power level:", power)
		var level *Level = NewLevel(levelStr, power)
		for fullRound, elfDied := level.PlayRound(); fullRound && !elfDied; {
			if debug {
				fmt.Println(level)
			}
			fullRound, elfDied = level.PlayRound()
			if !fullRound && !elfDied {
				elfDead = false
				fmt.Println(level)
				fmt.Println(level.Score())
			}
		}
	}
}

func main() {
	//part1(input, false)
	part2(input, false)
}
