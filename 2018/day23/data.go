package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func getInput() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(content))
}

func ParseStr(str string) []*Bot {
	re := regexp.MustCompile(`pos=<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>, r=([0-9]+)`)
	strs := strings.Split(str, "\n")
	bots := make([]*Bot, 0, len(strs))
	for _, s := range strings.Split(str, "\n") {
		match := re.FindStringSubmatch(s)
		if len(match) == 0 {
			continue
		}
		bots = append(bots, &Bot{
			Point{atoi(match[1]), atoi(match[2]), atoi(match[3])},
			atoi(match[4])})
	}
	return bots
}

var inputStr string = getInput()

const testStr = `pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1
`
