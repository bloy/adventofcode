package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
)

func getInput() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(content)
}

func ParseStr(str string) *Base {
	var b *Base = NewBase()
	var cur *Room
	stack := list.New()
	for _, ch := range str {
		switch ch {
		case '^':
			cur = b.StartRoom
		case 'N':
			cur = b.AddRoom(NORTH, cur)
		case 'E':
			cur = b.AddRoom(EAST, cur)
		case 'S':
			cur = b.AddRoom(SOUTH, cur)
		case 'W':
			cur = b.AddRoom(WEST, cur)
		case '(':
			stack.PushFront(cur)
		case ')':
			stack.Remove(stack.Front())
		case '|':
			cur = stack.Front().Value.(*Room)
		case '$':
			fmt.Println("Done with string")
		}
	}
	return b
}

var inputStr string = getInput()

const (
	testStr1  string = `^WNE$`
	testStr2a string = `^N(E|W)N$`
	testStr2  string = `^ENWWW(NEEE|SSE(EE|N))$`
	testStr3  string = `^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$`
	testStr4  string = `^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$`
	testStr5  string = `^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$`
)
