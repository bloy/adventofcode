package main

import (
	"fmt"
	"strings"
)

type direction rune

const (
	up    direction = 'U'
	down  direction = 'D'
	right direction = 'R'
	left  direction = 'L'
)

type keypad map[string]map[direction]string

var squarePad = keypad{
	"1": {up: "1", down: "4", left: "1", right: "2"},
	"2": {up: "2", down: "5", left: "1", right: "3"},
	"3": {up: "3", down: "6", left: "2", right: "3"},
	"4": {up: "1", down: "7", left: "4", right: "5"},
	"5": {up: "2", down: "8", left: "4", right: "6"},
	"6": {up: "3", down: "9", left: "5", right: "6"},
	"7": {up: "4", down: "7", left: "7", right: "8"},
	"8": {up: "5", down: "8", left: "7", right: "9"},
	"9": {up: "6", down: "9", left: "8", right: "9"},
}

var diamondPad = keypad{
	"1": {up: "1", down: "3", left: "1", right: "1"},
	"2": {up: "2", down: "6", left: "2", right: "3"},
	"3": {up: "1", down: "7", left: "2", right: "4"},
	"4": {up: "4", down: "8", left: "3", right: "4"},
	"5": {up: "5", down: "5", left: "5", right: "6"},
	"6": {up: "2", down: "A", left: "5", right: "7"},
	"7": {up: "3", down: "B", left: "6", right: "8"},
	"8": {up: "4", down: "C", left: "7", right: "9"},
	"9": {up: "9", down: "9", left: "8", right: "9"},
	"A": {up: "6", down: "A", left: "A", right: "B"},
	"B": {up: "7", down: "D", left: "A", right: "C"},
	"C": {up: "8", down: "C", left: "B", right: "C"},
	"D": {up: "B", down: "D", left: "D", right: "D"},
}

type dataType [][]direction

func readInput(str string) (data dataType) {
	lines := strings.Split(str, "\n")
	data = make(dataType, len(lines))
	for i, line := range lines {
		data[i] = make([]direction, len(line))
		for j, c := range line {
			data[i][j] = direction(c)
		}
	}
	return data
}

func solve(data dataType, pad keypad) string {
	value := ""
	pos := "5"
	for _, line := range data {
		for _, dir := range line {
			pos = pad[pos][dir]
			fmt.Println(string(dir), pos)
		}
		value = value + pos
	}
	return value
}

func solve1(data dataType) string {
	return ""
	//return solve(data, squarePad)
}

func solve2(data dataType) string {
	return solve(data, diamondPad)
}
