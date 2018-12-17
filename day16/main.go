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
	var codeMap map[int]OpcodeFunc = make(map[int]OpcodeFunc, len(gather))
	var possibles map[int]*list.List = make(map[int]*list.List)
	for code := range gather {
		fmt.Printf("code: %d, samples: %d\n", code, len(gather[code]))
		possibles[code] = list.New()
		for _, f := range funcs {
			matched := true
			for _, s := range gather[code] {
				if !matchSample(f, s) {
					matched = false
					break
				}
			}
			if matched {
				fmt.Printf(" found! - %v\n", f)
				possibles[code].PushBack(f)
			}
		}
		if possibles[code].Len() == 1 {
			codeMap[code] = possibles[code].Front().Value.(OpcodeFunc)
			delete(possibles, code)
		}
	}
	// TODO: while the length of codeMap < length of gather, look for a
	// possible with a len of 1. Add it to codemap, remove it from possibles,
	// repeat until done
	// TODO: parse code
	// TODO: run code
}
