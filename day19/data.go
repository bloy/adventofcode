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

func ParseProgram(programStr string) (ipReg int, instructions []Instruction) {
	ipReg = 0
	lines := strings.Split(programStr, "\n")
	instructions = make([]Instruction, 0, len(lines))
	for _, line := range lines {
		nums := strings.Split(line, " ")
		if line[0] == '#' {
			ipReg = atoi(nums[1])
		} else {
			instr := Instruction{}
			instr.Opcode = nums[0]
			instr.Input1 = atoi(nums[1])
			instr.Input2 = atoi(nums[2])
			instr.Output = atoi(nums[3])
			instructions = append(instructions, instr)
		}
	}
	return
}

const testStr = `#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`

const inputStr = `#ip 5
addi 5 16 5
seti 1 8 4
seti 1 5 3
mulr 4 3 1
eqrr 1 2 1
addr 1 5 5
addi 5 1 5
addr 4 0 0
addi 3 1 3
gtrr 3 2 1
addr 5 1 5
seti 2 5 5
addi 4 1 4
gtrr 4 2 1
addr 1 5 5
seti 1 2 5
mulr 5 5 5
addi 2 2 2
mulr 2 2 2
mulr 5 2 2
muli 2 11 2
addi 1 8 1
mulr 1 5 1
addi 1 18 1
addr 2 1 2
addr 5 0 5
seti 0 7 5
setr 5 0 1
mulr 1 5 1
addr 5 1 1
mulr 5 1 1
muli 1 14 1
mulr 1 5 1
addr 2 1 2
seti 0 0 0
seti 0 9 5`
