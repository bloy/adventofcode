package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IntcodeFunc is an intcode processiong emulator instruction
type IntcodeFunc func(ic *Intcode, positions ...int) (done bool, err error)

type intcodeopcode struct {
	args   int
	icfunc IntcodeFunc
}

// Intcode holds data and state for a running AoC 2019 intcode simulator
type Intcode struct {
	pc      int // program counter
	mem     []int
	opcodes map[int]intcodeopcode
}

// NewIntcode returns a new Intcode computer
func NewIntcode(codes []int) *Intcode {
	copied := make([]int, len(codes))
	for i := range codes {
		copied[i] = codes[i]
	}
	return &Intcode{pc: 0, mem: copied, opcodes: make(map[int]intcodeopcode)}
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

// AddOpcode adds an opcode to the existing opcodes
func (ic *Intcode) AddOpcode(opCodeNum, numArgs int, icfunc IntcodeFunc) {
	ic.opcodes[opCodeNum] = intcodeopcode{args: numArgs, icfunc: icfunc}
}

// RunInstruction Runs the single instruction at the program counter and
// updates the program counter
func (ic *Intcode) RunInstruction() (done bool, err error) {
	opcodenum := ic.mem[ic.pc]
	opcode, ok := ic.opcodes[opcodenum]
	if !ok {
		return true, fmt.Errorf(
			"Unknown opcode %d encountered at position %d", opcodenum, ic.pc,
		)
	}
	args := make([]int, opcode.args)
	for i := 1; i <= opcode.args; i++ {
		args[i-1] = ic.mem[ic.pc+i]
	}
	return opcode.icfunc(ic, args...)
}
