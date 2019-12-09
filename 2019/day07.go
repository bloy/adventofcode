package main

import (
	"bufio"
)

func init() {
	AddSolution(7, solveDay7)

}

func permNumber(perm []int64) int64 {
	var num int64
	for _, v := range perm {
		num = num*10 + v
	}
	return num
}

func permutations(start, stop int64) (permList [][]int64) {
	base := make([]int64, stop-start+1)
	for i := 0; i < len(base); i++ {
		j := int64(i)
		base[j] = j + start
	}
	var rc func([]int64, int64)
	rc = func(a []int64, k int64) {
		if k == int64(len(a)) {
			permList = append(permList, append([]int64{}, a...))
		} else {
			for i := k; i < int64(len(base)); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(base, 0)
	return permList
}

func solveDay7(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	program := ""
	for scanner.Scan() {
		program = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	permList1 := permutations(0, 4)
	permList2 := permutations(5, 9)
	pr.ReportLoad()

	var max, maxPerm int64

	for _, perm := range permList1 {
		signalchans := make([]chan int64, len(perm)+1)
		for i := 0; i < len(signalchans); i++ {
			signalchans[i] = make(chan int64, 2)
		}
		errorchan := make(chan error)
		donechan := make(chan bool)
		amps := make([]*Intcode, len(perm))
		for i, signal := range perm {
			amp, err := NewIntcodeFromInput(program)
			if err != nil {
				pr.logger.Fatal(err)
			}
			amps[i] = amp
			amps[i].AddStandardOpcodes()
			amps[i].RunProgramChannelMode(signalchans[i], signalchans[i+1], errorchan, donechan)
			signalchans[i] <- signal
		}
		signalchans[0] <- 0
		for i := 0; i < len(perm); i++ {
			<-donechan
		}
		out := <-signalchans[len(perm)]
		if out > max {
			max = out
			maxPerm = permNumber(perm)
		}
	}
	pr.ReportPart("Part1: Signal:", max, "Phase:", maxPerm)

	max = 0
	maxPerm = 0
	for _, perm := range permList2 {
		//fmt.Println("Trying perm", perm)
		signalchans := make([]chan int64, len(perm))
		for i := 0; i < len(signalchans); i++ {
			signalchans[i] = make(chan int64, 2)
		}
		errorchan := make(chan error)
		donechan := make(chan bool)
		amps := make([]*Intcode, len(perm))
		for i, signal := range perm {
			amp, err := NewIntcodeFromInput(program)
			if err != nil {
				pr.logger.Fatal(err)
			}
			amps[i] = amp
			amps[i].AddStandardOpcodes()
			inchan := signalchans[i]
			outchan := signalchans[0]
			if i != len(perm)-1 {
				outchan = signalchans[i+1]
			}
			amps[i].RunProgramChannelMode(inchan, outchan, errorchan, donechan)
			signalchans[i] <- signal
		}
		signalchans[0] <- 0
		for i := 0; i < len(perm); i++ {
			<-donechan
		}
		out := <-signalchans[0]
		if out > max {
			max = out
			maxPerm = permNumber(perm)
		}
	}

	pr.ReportPart("Part2: Signal:", max, "Phase:", maxPerm)
}
