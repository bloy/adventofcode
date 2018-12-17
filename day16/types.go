package main

type Registers [4]int

type Instruction struct {
	Opcode int
	Input1 int
	Input2 int
	Output int
}

type Sample struct {
	Before Registers
	Instruction
	After Registers
}

type OpcodeFunc func(int1, in2, out int, registers *Registers)
