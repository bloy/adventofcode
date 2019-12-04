package main

import (
	"bufio"
	"strconv"
	"strings"
)

func init() {
	AddSolution(4, solveDay4)
}

func day4ValidPassword(num int) bool {
	digits := make([]int, 6)
	n := num
	for i := 0; i < 6; i++ {
		digits[5-i] = n % 10
		n = n / 10
	}
	double := false
	increasing := true
	for i := 1; i < 6; i++ {
		if digits[i] == digits[i-1] {
			double = true
		} else if digits[i] < digits[i-1] {
			increasing = false
		}
	}
	return double && increasing
}

func day4ValidPart2(num int) bool {
	digits := make([]int, 6)
	n := num
	for i := 0; i < 6; i++ {
		digits[5-i] = n % 10
		n = n / 10
	}
	double := false
	for i := 1; i < 6; i++ {
		if digits[i] == digits[i-1] {
			count := 0
			for j := 0; j < 6; j++ {
				if digits[j] == digits[i] {
					count++
				}
			}
			if count == 2 {
				double = true
			}
		} else if digits[i] < digits[i-1] {
			return false
		}
	}
	return double
}

func solveDay4(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	top := 0
	bottom := 0
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "-")
		var err error
		top, err = strconv.Atoi(strs[1])
		if err != nil {
			pr.logger.Fatal(err)
		}
		bottom, err = strconv.Atoi(strs[0])
		if err != nil {
			pr.logger.Fatal(err)
		}
	}
	pr.ReportLoad()

	count := 0
	for i := bottom; i <= top; i++ {
		if day4ValidPassword(i) {
			count++
		}
	}
	pr.ReportPart(count, "total valid passwords")

	count = 0
	for i := bottom; i <= top; i++ {
		if day4ValidPart2(i) {
			count++
		}
	}
	pr.ReportPart(count, "total valid passwords")
}
