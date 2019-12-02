package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	AddSolution(2, solveDay2)
}

func solveDay2(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	var ints []int
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ",")
		ints = make([]int, len(strs))
		for i, s := range strs {
			num, err := strconv.Atoi(s)
			if err != nil {
				pr.logger.Fatal(err)
			}
			ints[i] = num
		}
	}
	orig := make([]int, len(ints))
	for i := range ints {
		orig[i] = ints[i]
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	pr.ReportLoad()

	pos := 0
	ints[1] = 12
	ints[2] = 2
	for ints[pos] != 99 {
		opcode := ints[pos]
		rhs1 := ints[pos+1]
		rhs2 := ints[pos+2]
		lhs := ints[pos+3]
		switch opcode {
		case 99:
			pr.logger.Println(pos, ints[pos], "HALT")
			break
		case 1:
			ints[lhs] = ints[rhs1] + ints[rhs2]
		case 2:
			ints[lhs] = ints[rhs1] * ints[rhs2]
		default:
			pr.logger.Fatalf("Unknown opcode %d at position %d", ints[pos], pos)
		}
		pos += 4
	}

	pr.ReportPart("Value at position 0:", ints[0])
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			ints = make([]int, len(orig))
			for i := range orig {
				ints[i] = orig[i]
			}
			ints[1] = n
			ints[2] = v
			pos := 0
			for ints[pos] != 99 {
				opcode := ints[pos]
				rhs1 := ints[pos+1]
				rhs2 := ints[pos+2]
				lhs := ints[pos+3]
				switch opcode {
				case 99:
					pr.logger.Println(pos, ints[pos], "HALT")
					break
				case 1:
					ints[lhs] = ints[rhs1] + ints[rhs2]
				case 2:
					ints[lhs] = ints[rhs1] * ints[rhs2]
				default:
					pr.logger.Fatalf("Unknown opcode %d at position %d", ints[pos], pos)
				}
				pos += 4
			}
			if ints[0] == 19690720 {
				pr.ReportPart("Part 2 SUCCESS", fmt.Sprintf("%02d%02d", n, v))
				return
			}
		}
	}
	pr.ReportPart("Part 2 FAILURE")
}
