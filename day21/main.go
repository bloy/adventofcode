package main

import "fmt"

func doPart1(str string) {
	ipReg, instructions := ParseProgram(str)
	minInstCount := 9999
	regValue := 0
	// this value was there as a check after analizing the code for when register 0 is used
	for i := 3909248; ; i++ {
		var registers Registers = Registers{}
		registers[0] = i
		var ip int = 0
		var instCount = 0
		for ip < len(instructions) && ip >= 0 && instCount < 10000 {
			instCount++
			registers[ipReg] = ip
			instr := instructions[ip]
			if ip == 28 {
				fmt.Println(instr, registers)
			}
			f := FuncNames[instr.Opcode]
			f(instr.Input1, instr.Input2, instr.Output, &registers)
			ip = registers[ipReg]
			ip++
		}
		if minInstCount > instCount {
			minInstCount = instCount
			regValue = i
		}
		if minInstCount < 9999 {
			break
		}
	}
	fmt.Println("part 1 min count:", minInstCount, "reg value:", regValue)
}

func doPart2(str string) {
	ipReg, instructions := ParseProgram(str)
	// this value was there as a check after analizing the code for when register 0 is used
	var registers Registers = Registers{}
	var ip int = 0
	var last = 0
	seen := make(map[int]bool)
	for ip < len(instructions) && ip >= 0 {
		registers[ipReg] = ip
		instr := instructions[ip]
		f := FuncNames[instr.Opcode]
		f(instr.Input1, instr.Input2, instr.Output, &registers)
		ip = registers[ipReg] + 1
		if registers[ipReg] == 28 {
			if seen[registers[3]] {
				break
			}
			seen[registers[3]] = true
			last = registers[3]
		}
	}
	fmt.Println("part 2:", last)
}

func main() {
	doPart1(inputStr)
	doPart2(inputStr)
}
