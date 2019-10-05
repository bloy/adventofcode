package main

type Registers [6]int

type Instruction struct {
	Opcode string
	Input1 int
	Input2 int
	Output int
}

type OpcodeFunc func(int1, in2, out int, registers *Registers)
