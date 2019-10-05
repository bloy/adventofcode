package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

type ItemStack struct {
	items []string
	lock  sync.RWMutex
}

func (s *ItemStack) New() *ItemStack {
	s.items = []string{}
	return s
}

func (s *ItemStack) Push(i string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, i)
}

func (s *ItemStack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *ItemStack) Empty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items) == 0
}

func getInput() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(content)
	//str = `dabAcCaCBAcCcaDA`
	str = strings.TrimSpace(str)
	return str
}

func reactPolymer(polymer string) string {
	var stack ItemStack
	stack.New()
	for i := 0; i < len(polymer); i++ {
		s1 := string(polymer[i])
		if stack.Empty() {
			stack.Push(s1)
		} else {
			s2 := stack.Pop()
			if s1 == s2 || strings.ToUpper(s1) != strings.ToUpper(s2) {
				stack.Push(s2)
				stack.Push(s1)
			}
		}
	}
	var build strings.Builder
	for _, ch := range stack.items {
		build.WriteString(ch)
	}
	return build.String()
}

func runPart1(input string) {
	fmt.Println("initial len", len(input))
	output := reactPolymer(input)
	fmt.Println("part 1", len(output))
}

func runPart2(input string) {
	var shortest int = len(input)
	for ch := 'a'; ch <= 'z'; ch++ {
		unit1 := string(ch)
		unit2 := strings.ToUpper(unit1)
		polymer := strings.Replace(strings.Replace(input, unit1, "", -1), unit2, "", -1)
		polymer = reactPolymer(polymer)
		pLen := len(polymer)
		if pLen < shortest {
			shortest = pLen
		}
	}
	fmt.Println("part 2", shortest)
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
