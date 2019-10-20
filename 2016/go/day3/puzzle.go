package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type dataType [][]int

func validTriangle(s1, s2, s3 int) bool {
	return (s1 < s2+s3 && s2 < s1+s3 && s3 < s1+s2)
}

func readInput(str string) (data dataType) {
	spaces := regexp.MustCompile(`\s+`)
	lines := strings.Split(str, "\n")
	data = make([][]int, len(lines))
	for i, line := range lines {
		numStrs := spaces.Split(line, -1)
		data[i] = make([]int, len(numStrs))
		for j, s := range numStrs {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Number conversion error:", err)
			}
			data[i][j] = n
		}
	}
	return
}

func solve1(data dataType) int {
	var count int
	for _, row := range data {
		if validTriangle(row[0], row[1], row[2]) {
			count++
		}
	}
	return count
}

func solve2(data dataType) int {
	var count int
	for i := 0; i < len(data); i = i + 3 {
		for j := 0; j < 3; j++ {
			if validTriangle(data[i][j], data[i+1][j], data[i+2][j]) {
				count++
			}
		}
	}
	return count
}
