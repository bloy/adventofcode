package main

import (
	"container/ring"
	"fmt"
)

func runPart1(playerCount, maxMarble int) {
	scores := make([]int, playerCount)
	current := ring.New(1)
	current.Value = 0
	player := 0
	for marble := 1; marble <= maxMarble; marble++ {
		if marble%23 == 0 {
			scores[player] += marble
			removed := current
			for i := 0; i < 7; i++ {
				removed = removed.Prev()
			}
			scores[player] += removed.Value.(int)
			current = removed.Next()
			removed.Prev().Link(removed.Next())
		} else {
			after := current.Next()
			insert := ring.New(1)
			insert.Value = marble
			after.Link(insert)
			current = insert
		}
		//fmt.Printf("[%3d] %d", player, current.Value)
		//for node := current.Next(); node != current; node = node.Next() {
		//fmt.Printf(" %d", node.Value)
		//}
		//fmt.Println("")
		player = (player + 1) % playerCount
	}
	max := 0
	for i := range scores {
		if scores[i] > max {
			max = scores[i]
		}
	}
	fmt.Println("max score:", max)
}

func main() {
	const playerCount int = 403
	const maxMarble int = 71920
	//const playerCount int = 10
	//const maxMarble int = 1618
	runPart1(playerCount, maxMarble)
	runPart1(playerCount, 100*maxMarble)
}
