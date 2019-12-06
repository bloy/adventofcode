package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IntcodeFunc is an intcode processiong emulator instruction
type IntcodeFunc func(ic *Intcode, params []int, positions []int) (done bool, err error)

type intcodeopcode struct {
	args   int
	icfunc IntcodeFunc
}

func opcodeHalt(ic *Intcode, params []int, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("HALT")
	}
	ic.pc++
	return true, nil
}

func opcodeAdd(ic *Intcode, params []int, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("ADD ", positions)
	}
	i1 := positions[0]
	i2 := positions[1]
	if params[0] == 0 {
		i1 = ic.mem[i1]
	}
	if params[1] == 0 {
		i2 = ic.mem[i2]
	}
	o := positions[2]
	ic.mem[o] = i1 + i2
	ic.pc += 4
	return
}

func opcodeMul(ic *Intcode, params []int, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("MUL ", positions)
	}
	i1 := positions[0]
	i2 := positions[1]
	if params[0] == 0 {
		i1 = ic.mem[i1]
	}
	if params[1] == 0 {
		i2 = ic.mem[i2]
	}
	o := positions[2]
	ic.mem[o] = i1 * i2
	ic.pc += 4
	return
}

func opcodeInput(ic *Intcode, params []int, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("IN  ", positions)
	}
	if len(ic.inputs) < 1 {
		return true, fmt.Errorf("No Inputs Remaining")
	}
	in := ic.inputs[0]
	ic.inputs = ic.inputs[1:]
	o := positions[0]
	ic.mem[o] = in
	ic.pc += 2
	return
}

func opcodeOutput(ic *Intcode, params []int, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("OUT ", positions)
	}
	in := positions[0]
	if params[0] == 0 {
		in = ic.mem[in]
	}
	ic.output = append(ic.output, in)
	ic.pc += 2
	return
}

func opcodeJumpIfTrue(ic *Intcode, params []int, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("JMPT", positions)
	}
	in := make([]int, 2)
	for i := 0; i < 2; i++ {
		in[i] = positions[i]
		if params[i] == 0 {
			in[i] = ic.mem[in[i]]
		}
	}
	if in[0] != 0 {
		ic.pc = in[1]
	} else {
		ic.pc++
	}
	return
}

func opcodeJumpIfFalse(ic *Intcode, params []int, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("JMPF", positions)
	}
	in := make([]int, 2)
	for i := 0; i < 2; i++ {
		in[i] = positions[i]
		if params[i] == 0 {
			in[i] = ic.mem[in[i]]
		}
	}
	if in[0] == 0 {
		ic.pc = in[1]
	} else {
		ic.pc++
	}
	return
}

func opcodeLessThan(ic *Intcode, params []int, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("LT  ", positions)
	}
	in := make([]int, 3)
	for i := 0; i < 3; i++ {
		in[i] = positions[i]
		if params[i] == 0 && i != 2 {
			in[i] = ic.mem[in[i]]
		}
	}
	out := 0
	if in[0] < in[1] {
		out = 1
	}
	ic.mem[in[2]] = out
	ic.pc++
	return
}

func opcodeEqual(ic *Intcode, params []int, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("EQ  ", positions)
	}
	in := make([]int, 3)
	for i := 0; i < 3; i++ {
		in[i] = positions[i]
		if params[i] == 0 && i != 2 {
			in[i] = ic.mem[in[i]]
		}
	}
	out := 0
	if in[0] == in[1] {
		out = 1
	}
	ic.mem[in[2]] = out
	ic.pc++
	return
}

// Intcode holds data and state for a running AoC 2019 intcode simulator
type Intcode struct {
	pc      int // program counter
	mem     []int
	opcodes map[int]intcodeopcode
	output  []int
	inputs  []int
	Verbose bool
}

// NewIntcode returns a new Intcode computer
func NewIntcode(codes []int) *Intcode {
	copied := make([]int, len(codes))
	for i := range codes {
		copied[i] = codes[i]
	}
	return &Intcode{
		pc:      0,
		mem:     copied,
		opcodes: make(map[int]intcodeopcode),
		output:  make([]int, 0),
		inputs:  make([]int, 0),
	}
}

// NewIntcodeFromInput returns a new Intcode built from a string of comma
// seperated integers
func NewIntcodeFromInput(codes string) (*Intcode, error) {
	strs := strings.Split(codes, ",")
	nums := make([]int, len(strs))
	for i := range strs {
		num, err := strconv.Atoi(strs[i])
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}
	return NewIntcode(nums), nil
}

// AddStandardOpcodes adds the standard opcodes
func (ic *Intcode) AddStandardOpcodes() {
	ic.AddOpcode(99, 0, opcodeHalt)       // HALT
	ic.AddOpcode(1, 3, opcodeAdd)         // ADD
	ic.AddOpcode(2, 3, opcodeMul)         // MUL
	ic.AddOpcode(3, 1, opcodeInput)       // IN
	ic.AddOpcode(4, 1, opcodeOutput)      // OUT
	ic.AddOpcode(5, 2, opcodeJumpIfTrue)  // JMPT
	ic.AddOpcode(6, 2, opcodeJumpIfFalse) // JMPF
	ic.AddOpcode(7, 3, opcodeLessThan)    // LT
	ic.AddOpcode(8, 3, opcodeEqual)       // EQ
}

// AddOpcode adds an opcode to the existing opcodes
func (ic *Intcode) AddOpcode(opCodeNum, numArgs int, icfunc IntcodeFunc) {
	ic.opcodes[opCodeNum] = intcodeopcode{args: numArgs, icfunc: icfunc}
}

// RunInstruction Runs the single instruction at the program counter and
// updates the program counter
func (ic *Intcode) RunInstruction() (done bool, err error) {
	ocflags := ic.mem[ic.pc]
	opcodenum := ocflags % 100
	ocflags = ocflags / 100
	opcode, ok := ic.opcodes[opcodenum]
	if !ok {
		return true, fmt.Errorf(
			"Unknown opcode %d encountered at position %d", opcodenum, ic.pc,
		)
	}
	args := make([]int, opcode.args)
	flags := make([]int, opcode.args)
	for i := 1; i <= opcode.args; i++ {
		flag := ocflags % 10
		ocflags = ocflags / 10
		flags[i-1] = flag
		arg := ic.mem[ic.pc+i]
		args[i-1] = arg
	}
	done, err = opcode.icfunc(ic, flags, args)
	if ic.Verbose {
		fmt.Println("   ", ic.pc, ic.mem)
	}
	return
}

// RunProgram is a method on *IntCode
func (ic *Intcode) RunProgram(inputs []int) (output []int, err error) {
	if inputs == nil {
		inputs = make([]int, 0)
	}
	ic.inputs = inputs
	var done bool
	for !done && err == nil {
		done, err = ic.RunInstruction()
	}
	return ic.output, err
}
