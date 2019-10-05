package main

import "container/list"
import "fmt"

type countNode struct {
	count int
	room  *Room
}

func doPart1(str string) {
	b := ParseStr(str)
	queue := list.New()
	seen := make(map[Point]bool)
	cur := b.StartRoom
	seen[cur.Position] = true
	maxLen := 0
	queue.PushBack(countNode{0, cur})
	for queue.Front() != nil {
		node := queue.Front().Value.(countNode)
		queue.Remove(queue.Front())
		if node.count > maxLen {
			maxLen = node.count
		}
		for _, v := range node.room.Doors {
			if v != nil {
				if !seen[v.Position] {
					seen[v.Position] = true
					queue.PushBack(countNode{node.count + 1, v})
				}
			}
		}
	}
	fmt.Println("part 1", maxLen)
}

func pathLenTo(b *Base, target *Room) int {
	queue := list.New()
	seen := make(map[Point]bool)
	cur := b.StartRoom
	seen[cur.Position] = true
	queue.PushBack(countNode{0, cur})
	for queue.Front() != nil {
		node := queue.Front().Value.(countNode)
		queue.Remove(queue.Front())
		if node.room.Position == target.Position {
			return node.count
		}
		for _, d := range node.room.Doors {
			if d != nil && !seen[d.Position] {
				seen[d.Position] = true
				queue.PushBack(countNode{node.count + 1, d})
			}
		}
	}
	return -1
}

func doPart2(str string) {
	b := ParseStr(str)
	farRooms := 0
	for _, r := range b.Rooms {
		if pathLenTo(b, r) >= 1000 {
			farRooms++
		}
	}
	fmt.Println("part2:", farRooms)
}

func main() {
	doPart1(inputStr)
	doPart2(inputStr)
}
