package main

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func init() {
	AddSolution(2, solveDay2)
}

func solveDay2(pr *PuzzleRun) {
	s := bufio.NewScanner(pr.InFile)
	boxes := make([][]int, 0)

	for s.Scan() {
		strs := strings.Split(s.Text(), "x")
		box := make([]int, 3)
		for i := 0; i < 3; i++ {
			num, err := strconv.Atoi(strs[i])
			if err != nil {
				pr.logger.Fatal(err)
			}
			box[i] = num
		}
		sort.Slice(box, func(i, j int) bool { return box[i] < box[j] })
		boxes = append(boxes, box)
	}
	if err := s.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	var total int

	for _, box := range boxes {
		total += box[0]*box[1]*3 + box[0]*box[2]*2 + box[1]*box[2]*2
	}
	pr.ReportPart("Wrapping Paper:", total, "square feet")

	total = 0

	for _, box := range boxes {
		total += box[0]*2 + box[1]*2 + box[0]*box[1]*box[2]
	}
	pr.ReportPart("Ribbon:", total, "feet")
}
