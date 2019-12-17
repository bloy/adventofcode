package main

import (
	"bufio"
	"strconv"
)

func init() {
	AddSolution(16, solveDay16)
}

var fftBasePattern = []int{0, 1, 0, -1}

func day16FFT(in []int) (out []int) {
	out = make([]int, len(in))
	pattern := fftBasePattern
	for i := range out {
		sum := 0
		for j := range in {
			mul := pattern[(j+1)/(i+1)%4]
			sum += mul * in[j]
		}
		if sum < 0 {
			sum *= -1
		}
		out[i] = sum % 10
	}
	return
}

func solveDay16(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	nums := []int{}
	for scanner.Scan() {
		str := scanner.Text()
		for i := range str {
			num, err := strconv.Atoi(string(str[i]))
			pr.CheckError(err)
			nums = append(nums, num)
		}
	}
	if err := scanner.Err(); err != nil {
		pr.CheckError(err)
	}
	pr.ReportLoad()

	val := nums
	for i := 0; i < 100; i++ {
		val = day16FFT(val)
	}
	pr.ReportPart(val[:8])

	offsetSlice := nums[:7]
	offset := 0
	for i := 0; i < len(offsetSlice); i++ {
		offset = offset*10 + offsetSlice[i]
	}
	pr.ReportPart(len(nums), offset, len(nums)*10000-offset)
}
