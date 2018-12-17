package main

import (
	"fmt"
)

func main() {
	samples := ParseSamples(SampleStr)
	doPart1(samples)

}

func doPart1(samples []*Sample) {
	var sampleCount int
	for _, sample := range samples {
		var count int = 0
		for _, f := range funcs {
			tmp := regCopy(sample.Before)
			f(sample.Input1, sample.Input2, sample.Output, &tmp)
			if tmp == sample.After {
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
