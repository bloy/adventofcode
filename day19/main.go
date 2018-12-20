package main

import "fmt"

func doPart1(ipReg int, instructions []Instruction) {
	var debug bool = false
	var ip int
	var registers Registers
	for ip < len(instructions) && ip >= 0 {
		if debug {
			fmt.Printf("ip=%d %v ", ip, registers)
		}
		registers[ipReg] = ip
		instr := instructions[ip]
		if debug {
			fmt.Printf("%v ", instr)
		}
		f := FuncNames[instr.Opcode]
		f(instr.Input1, instr.Input2, instr.Output, &registers)
		ip = registers[ipReg]
		ip++
		if debug {
			fmt.Printf("%v\n", registers)
		}
	}
	fmt.Println("part1", registers)
}

func doPart2(ipReg int, instructions []Instruction) {
	var debug bool = false
	var ip int
	var registers Registers
	registers[0] = 1
	for ip < len(instructions) && ip >= 0 {
		if ip == 1 {
			// once ip == 1, the program has calculated the number it wants
			break
		}
		debug = (ip < 3 || ip > 11)
		if debug {
			fmt.Printf("ip=%2d ", ip)
		}
		registers[ipReg] = ip
		instr := instructions[ip]
		if debug {
			fmt.Printf("%v ", instr)
		}
		f := FuncNames[instr.Opcode]
		f(instr.Input1, instr.Input2, instr.Output, &registers)
		ip = registers[ipReg]
		ip++
		if debug {
			fmt.Printf("%v\n", registers)
		}
	}
	// optimize the assembly to avoid an exponential number of muls
	num := registers[2]
	sum := 0
	for i := 1; i <= num; i++ {
		div := num / i
		if div*i == num {
			sum += i
		}
	}
	fmt.Println("part2", sum)
}

func main() {
	ipReg, instructions := ParseProgram(inputStr)
	doPart1(ipReg, instructions)
	doPart2(ipReg, instructions)
}
