package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	AddSolution(16, solveDay16)
}

var fftBasePattern = []int{0, 1, 0, -1}

func day16FFT(in []int, offset int) (out []int) {
	out = make([]int, len(in))
	pattern := fftBasePattern
	for i := range out {
		sum := 0
		for j := range in {
			mul := pattern[(j+1+offset)/(i+1)%4]
			sum += mul * in[j]
		}
		if sum < 0 {
			sum *= -1
		}
		out[i] = sum % 10
	}
	return
}

func day16signal(slice []int) string {
	b := strings.Builder{}
	for i := range slice {
		fmt.Fprintf(&b, "%d", slice[i])
	}
	return b.String()
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
		val = day16FFT(val, 0)
	}
	pr.ReportPart(day16signal(val[:8]))

	reps := 10000
	phases := 100
	offset := 0
	for i := 0; i < 7; i++ {
		offset = offset*10 + nums[i]
	}
	val = make([]int, (len(nums)*reps)-offset)
	for i := range val {
		val[i] = nums[(offset+i)%(len(nums))]
	}
	if offset < len(nums)*reps/2 {
		pr.logger.Fatal("offset is too small", offset)
	}

	fmt.Println(len(nums), len(val))

	// since we know offset > len(nums)*reps/2, all coefficients are 1
	for phase := phases; phase > 0; phase-- {
		sum := 0
		for i := len(val) - 1; i >= 0; i-- {
			sum += val[i]
			n := sum
			if n < 0 {
				n = n * -1
			}
			val[i] = n % 10
		}
	}
	pr.ReportPart(day16signal(val[:8]))
}
