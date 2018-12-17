package main

import (
	"strconv"
	"strings"
)

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func regCopy(reg Registers) (newReg Registers) {
	for i := 0; i < 4; i++ {
		newReg[i] = reg[i]
	}
	return
}

func ParseSamples(sampleStr string) (samples []*Sample) {
	lines := strings.Split(sampleStr, "\n")
	samples = make([]*Sample, 0, len(lines)/4)
	for i := 0; i < len(lines); i += 4 {
		var s *Sample
		s = &Sample{}
		beforeStrs := strings.Split(lines[i][9:len(lines[i])-1], ", ")
		afterStrs := strings.Split(lines[i+2][9:len(lines[i+2])-1], ", ")
		for j := 0; j < 4; j++ {
			s.Before[j] = atoi(beforeStrs[j])
			s.After[j] = atoi(afterStrs[j])
		}
		numStrs := strings.Split(lines[i+1], " ")
		s.Opcode = atoi(numStrs[0])
		s.Input1 = atoi(numStrs[1])
		s.Input2 = atoi(numStrs[2])
		s.Output = atoi(numStrs[3])
		samples = append(samples, s)
	}
	return
}
