package main

import "container/heap"
import "fmt"

func main() {
	depth := inputDepth
	target := inputTarget
	//depth := testDepth
	//target := testTarget
	fmt.Println(depth, target)
	cave := NewCave(depth, target.X, target.Y)
	fmt.Println(cave)
	totalRisk := 0
	for y := 0; y <= target.Y; y++ {
		for x := 0; x <= target.X; x++ {
			r := cave.Region(x, y)
			totalRisk += r.Risk()
		}
	}
	fmt.Println("part 1 total risk:", totalRisk)

	seen := make(map[SeenNode]int)
	pq := make(PriorityQueue, 1)
	pq[0] = &Node{Point{0, 0}, 0, TORCH}
	heap.Init(&pq)
	seen[SeenNode{Point{0, 0}, TORCH}] = 0
	var node *Node
	for pq.Len() > 0 {
		node = heap.Pop(&pq).(*Node)
		if node.Point == cave.Target && node.Equipped == TORCH {
			break // done here
		}

		for _, dir := range DIRECTIONS {
			newSeenNode := SeenNode{Point{node.X + dir.X, node.Y + dir.Y}, node.Equipped}
			if newSeenNode.X < 0 || newSeenNode.Y < 0 {
				continue
			}
			r := cave.Region(newSeenNode.X, newSeenNode.Y)
			if ItemNotAllowed[r.Type] == node.Equipped {
				continue
			}
			if seenTime, ok := seen[newSeenNode]; ok && seenTime <= node.Minutes+1 {
				continue
			}
			newNode := Node{Point{newSeenNode.X, newSeenNode.Y}, node.Minutes + 1, node.Equipped}
			heap.Push(&pq, &newNode)
			seen[newSeenNode] = node.Minutes + 1
		}
		var altEquip Item
		r := cave.Region(node.X, node.Y)
		for _, item := range EQUIPMENT {
			if item != node.Equipped && item != ItemNotAllowed[r.Type] {
				altEquip = item
			}
		}
		newSeenNode := SeenNode{Point{node.X, node.Y}, altEquip}
		if seenTime, ok := seen[newSeenNode]; ok && seenTime <= node.Minutes+7 {
			continue
		}
		newNode := Node{Point{node.X, node.Y}, node.Minutes + 7, altEquip}
		heap.Push(&pq, &newNode)
		seen[newSeenNode] = node.Minutes + 7
	}
	fmt.Println("part2:", node.Minutes, "minutes")
}
