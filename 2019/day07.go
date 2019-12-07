package main

import "bufio"

func init() {
	AddSolution(7, solveDay7)

}

func permNumber(perm []int) int {
	num := 0
	for _, v := range perm {
		num = num*10 + v
	}
	return num
}

func permutations(start, stop int) (permList [][]int) {
	base := make([]int, stop-start+1)
	for i := start; i < len(base); i++ {
		base[i-start] = i
	}
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permList = append(permList, append([]int{}, a...))
		} else {
			for i := k; i < len(base); i++ {
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
	pr.ReportLoad()

	max := 0
	maxPerm := 0

	for _, perm := range permList1 {
		out := 0
		for _, signal := range perm {
			ic, err := NewIntcodeFromInput(program)
			if err != nil {
				pr.logger.Fatal(err)
			}
			ic.AddStandardOpcodes()
			output, err := ic.RunProgram([]int{signal, out})
			if err != nil {
				pr.logger.Fatal(err)
			}
			out = output[0]
		}
		if out > max {
			max = out
			maxPerm = permNumber(perm)
		}
	}
	pr.ReportPart("Part1: Signal:", max, "Phase:", maxPerm)
}
