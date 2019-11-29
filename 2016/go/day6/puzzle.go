package main

import (
	"sort"
	"strings"
	"text/scanner"
)

type runeCounter map[rune]int

func readInput(inStr string) []string {
	var s scanner.Scanner
	var msgs []string
	s.Init(strings.NewReader(inStr))
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		msgs = append(msgs, s.TokenText())
	}
	return msgs
}

func solve1(data []string) string {
	var counters []runeCounter
	var message []rune
	var counterMax []int
	if len(data) == 0 {
		return ""
	}
	size := len(data[0])
	counters = make([]runeCounter, size)
	message = make([]rune, size)
	counterMax = make([]int, size)
	for i := 0; i < size; i++ {
		counters[i] = make(runeCounter)
	}
	for _, msg := range data {
		for i := 0; i < len(msg); i++ {
			c := rune(msg[i])
			count := counters[i][c] + 1
			counters[i][c] = count
			if count > counterMax[i] {
				counterMax[i] = count
				message[i] = c
			}
		}
	}
	return string(message)
}

func leastCommon(counter runeCounter) rune {
	runes := make([]rune, 0, len(counter))
	for r := range counter {
		runes = append(runes, r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return counter[runes[i]] < counter[runes[j]]
	})
	return runes[0]
}

func solve2(data []string) string {
	var counters []runeCounter
	size := len(data[0])
	counters = make([]runeCounter, size)
	for i := 0; i < size; i++ {
		counters[i] = make(runeCounter)
	}
	for _, msg := range data {
		for i := 0; i < len(msg); i++ {
			c := rune(msg[i])
			count := counters[i][c] + 1
			counters[i][c] = count
		}
	}
	message := make([]rune, size)
	for i := 0; i < size; i++ {
		message[i] = leastCommon(counters[i])
	}
	return string(message)
}
