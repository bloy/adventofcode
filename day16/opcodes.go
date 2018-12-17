package main

func addr(in1, in2, out int, registers *Registers) {
	registers[out] = registers[in1] + registers[in2]
}

func addi(in1, in2, out int, registers *Registers) {
	registers[out] = registers[in1] + in2
}

func mulr(in1, in2, out int, registers *Registers) {
	registers[out] = registers[in1] * registers[in2]
}

func muli(in1, in2, out int, registers *Registers) {
	registers[out] = registers[in1] * in2
}

func banr(in1, in2, out int, registers *Registers) {
	registers[out] = registers[in1] & registers[in2]
}

func bani(in1, in2, out int, registers *Registers) {
	registers[out] = registers[in1] & in2
}

func borr(in1, in2, out int, registers *Registers) {
	registers[out] = registers[in1] | registers[in2]
}

func bori(in1, in2, out int, registers *Registers) {
	registers[out] = registers[in1] | in2
}

func setr(in1, _, out int, registers *Registers) {
	registers[out] = registers[in1]
}

func seti(in1, _, out int, registers *Registers) {
	registers[out] = in1
}

func gtir(in1, in2, out int, registers *Registers) {
	if in1 > registers[in2] {
		registers[out] = 1
	} else {
		registers[out] = 0
	}
}

func gtri(in1, in2, out int, registers *Registers) {
	if registers[in1] > in2 {
		registers[out] = 1
	} else {
		registers[out] = 0
	}
}

func gtrr(in1, in2, out int, registers *Registers) {
	if registers[in1] > registers[in2] {
		registers[out] = 1
	} else {
		registers[out] = 0
	}
}

func eqir(in1, in2, out int, registers *Registers) {
	if in1 == registers[in2] {
		registers[out] = 1
	} else {
		registers[out] = 0
	}
}

func eqri(in1, in2, out int, registers *Registers) {
	if registers[in1] == in2 {
		registers[out] = 1
	} else {
		registers[out] = 0
	}
}

func eqrr(in1, in2, out int, registers *Registers) {
	if registers[in1] == registers[in2] {
		registers[out] = 1
	} else {
		registers[out] = 0
	}
}

var funcs [16]OpcodeFunc = [16]OpcodeFunc{
	addr, addi,
	mulr, muli,
	banr, bani,
	borr, bori,
	setr, seti,
	gtir, gtri, gtrr,
	eqir, eqri, eqrr,
}
