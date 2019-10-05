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

const inputStr string = `#ip 4
seti 123 0 3
bani 3 456 3
eqri 3 72 3
addr 3 4 4
seti 0 0 4
seti 0 2 3
bori 3 65536 2
seti 1397714 1 3
bani 2 255 5
addr 3 5 3
bani 3 16777215 3
muli 3 65899 3
bani 3 16777215 3
gtir 256 2 5
addr 5 4 4
addi 4 1 4
seti 27 6 4
seti 0 6 5
addi 5 1 1
muli 1 256 1
gtrr 1 2 1
addr 1 4 4
addi 4 1 4
seti 25 2 4
addi 5 1 5
seti 17 0 4
setr 5 7 2
seti 7 4 4
eqrr 3 0 5
addr 5 4 4
seti 5 8 4`
