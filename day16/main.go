package main

import (
	"container/list"
	"fmt"
)

func main() {
	samples := ParseSamples(SampleStr)
	doPart1(samples)
	doPart2(samples)
}

func matchSample(f OpcodeFunc, sample *Sample) bool {
	tmp := regCopy(sample.Before)
	f(sample.Input1, sample.Input2, sample.Output, &tmp)
	if tmp == sample.After {
		return true
	} else {
		return false
	}
}

func doPart1(samples []*Sample) {
	var sampleCount int
	for _, sample := range samples {
		var count int = 0
		for _, f := range funcs {
			if matchSample(f, sample) {
				count++
			}
			if count >= 3 {
				break
			}
		}
		if count >= 3 {
			sampleCount++
		}
	}
	fmt.Printf("%d of %d samples match 3 more more opcodes\n", sampleCount, len(samples))
}

func doPart2(samples []*Sample) {
	var gather map[int][]*Sample = make(map[int][]*Sample)
	for _, sample := range samples {
		code := sample.Opcode
		if gather[code] == nil {
			gather[code] = make([]*Sample, 0, 1)
		}
		gather[code] = append(gather[code], sample)
	}
	var codeMap map[int]OpcodeFunc = make(map[int]OpcodeFunc)
	var possibles map[int]*list.List = make(map[int]*list.List)
	for code := range gather {
		possibles[code] = list.New()
		for name, f := range FuncNames {
			matched := true
			for _, s := range gather[code] {
				if !matchSample(f, s) {
					matched = false
					break
				}
			}
			if matched {
				possibles[code].PushBack(name)
			}
		}
	}
	for len(codeMap) < len(gather) {
		possiblecodes := make([]int, 0)
		for code := range possibles {
			if possibles[code].Len() == 1 {
				possiblecodes = append(possiblecodes, code)
			}
		}
		for _, code := range possiblecodes {
			name := possibles[code].Front().Value.(string)
			f := FuncNames[name]
			codeMap[code] = f
			delete(possibles, code)
			for delcode := range possibles {
				node := possibles[delcode].Front()
				for node != nil {
					next := node.Next()
					if node.Value.(string) == name {
						possibles[delcode].Remove(node)
					}
					node = next
				}
			}
		}
	}
	program := ParseProgram(ProgramStr)
	var registers Registers
	for _, instr := range program {
		codeMap[instr.Opcode](instr.Input1, instr.Input2, instr.Output, &registers)
	}
	fmt.Println("after executing the proggram, register 0 contains: ", registers[0])
}
