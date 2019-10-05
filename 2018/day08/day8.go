package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const testStr string = `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`

type NodeType struct {
	children []NodeType
	metadata []int
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func readNode(strs []string) (node NodeType, remaining []string) {
	var numKids, numMeta int

	numKids = atoi(strs[0])
	numMeta = atoi(strs[1])
	node.children = make([]NodeType, numKids)
	node.metadata = make([]int, numMeta)
	remaining = strs[2:]
	for i := 0; i < numKids; i++ {
		node.children[i], remaining = readNode(remaining)
	}
	for i := 0; i < numMeta; i++ {
		node.metadata[i] = atoi(remaining[i])
	}
	remaining = remaining[numMeta:]
	return node, remaining
}

func getInput() NodeType {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(content)
	//str = testStr
	strs := strings.Split(strings.TrimSpace(str), " ")
	topNode, _ := readNode(strs)
	return topNode
}

func runPart1(input NodeType) {
	nodeStack := []NodeType{}
	var sum int
	nodeStack = append(nodeStack, input)
	for len(nodeStack) > 0 {
		node := nodeStack[0]
		nodeStack = nodeStack[1:]
		for _, meta := range node.metadata {
			sum += meta
		}
		for _, child := range node.children {
			nodeStack = append([]NodeType{child}, nodeStack...)
		}
	}

	fmt.Println("part 1", sum)
}

func nodeValue(node NodeType) int {
	var sum int
	for _, meta := range node.metadata {
		if len(node.children) == 0 {
			sum += meta
		} else {
			if meta > 0 && meta <= len(node.children) {
				sum += nodeValue(node.children[meta-1])
			}
		}
	}
	return sum
}

func runPart2(input NodeType) {
	fmt.Println("part 2", nodeValue(input))
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
