package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IntcodeFunc is an intcode processiong emulator instruction
type IntcodeFunc func(ic *Intcode, positions []int) (done bool, err error)

type intcodeopcode struct {
	args     int
	argflags string
	icfunc   IntcodeFunc
}

func opcodeHalt(ic *Intcode, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("HALT")
	}
	ic.pc++
	return true, nil
}

func opcodeAdd(ic *Intcode, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("ADD ", positions)
	}
	i1 := positions[0]
	i2 := positions[1]
	o := positions[2]
	ic.mem[o] = i1 + i2
	ic.pc += 4
	return
}

func opcodeMul(ic *Intcode, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("MUL ", positions)
	}
	i1 := positions[0]
	i2 := positions[1]
	o := positions[2]
	ic.mem[o] = i1 * i2
	ic.pc += 4
	return
}

func opcodeInput(ic *Intcode, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("IN  ", positions)
	}
	var in int
	if ic.UseChannels {
		in = <-ic.inchan
	} else {
		if len(ic.inputs) < 1 {
			return true, fmt.Errorf("No Inputs Remaining")
		}
		in = ic.inputs[0]
		ic.inputs = ic.inputs[1:]
	}
	o := positions[0]
	ic.mem[o] = in
	ic.pc += 2
	return
}

func opcodeOutput(ic *Intcode, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("OUT ", positions)
	}
	in := positions[0]
	if ic.UseChannels {
		ic.outchan <- in
	} else {
		ic.output = append(ic.output, in)
	}
	ic.pc += 2
	return
}

func opcodeJumpIfTrue(ic *Intcode, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("JMPT", positions)
	}
	in1 := positions[0]
	in2 := positions[1]
	if in1 != 0 {
		ic.pc = in2
	} else {
		ic.pc += 3
	}
	return
}

func opcodeJumpIfFalse(ic *Intcode, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("JMPF", positions)
	}
	in1 := positions[0]
	in2 := positions[1]
	if in1 == 0 {
		ic.pc = in2
	} else {
		ic.pc += 3
	}
	return
}

func opcodeLessThan(ic *Intcode, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("LT  ", positions)
	}
	in1 := positions[0]
	in2 := positions[1]
	o := positions[2]
	out := 0
	if in1 < in2 {
		out = 1
	}
	ic.mem[o] = out
	ic.pc += 4
	return
}

func opcodeEqual(ic *Intcode, positions []int) (done bool, err error) {
	if ic.Verbose {
		fmt.Println("EQ  ", positions)
	}
	in1 := positions[0]
	in2 := positions[1]
	o := positions[2]
	out := 0
	if in1 == in2 {
		out = 1
	}
	ic.mem[o] = out
	ic.pc += 4
	return
}

// Intcode holds data and state for a running AoC 2019 intcode simulator
type Intcode struct {
	pc          int // program counter
	mem         []int
	opcodes     map[int]intcodeopcode
	output      []int
	outchan     chan int
	inputs      []int
	inchan      chan int
	UseChannels bool
	Verbose     bool
}

// NewIntcode returns a new Intcode computer
func NewIntcode(codes []int) *Intcode {
	copied := make([]int, len(codes))
	for i := range codes {
		copied[i] = codes[i]
	}
	return &Intcode{
		pc:          0,
		mem:         copied,
		opcodes:     make(map[int]intcodeopcode),
		output:      make([]int, 0),
		inputs:      make([]int, 0),
		inchan:      make(chan int),
		outchan:     make(chan int),
		Verbose:     false,
		UseChannels: false,
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
	ic.AddOpcode(99, 0, "", opcodeHalt)         // HALT
	ic.AddOpcode(1, 3, "rrw", opcodeAdd)        // ADD
	ic.AddOpcode(2, 3, "rrw", opcodeMul)        // MUL
	ic.AddOpcode(3, 1, "w", opcodeInput)        // IN
	ic.AddOpcode(4, 1, "r", opcodeOutput)       // OUT
	ic.AddOpcode(5, 2, "rr", opcodeJumpIfTrue)  // JMPT
	ic.AddOpcode(6, 2, "rr", opcodeJumpIfFalse) // JMPF
	ic.AddOpcode(7, 3, "rrw", opcodeLessThan)   // LT
	ic.AddOpcode(8, 3, "rrw", opcodeEqual)      // EQ
}

// AddOpcode adds an opcode to the existing opcodes
func (ic *Intcode) AddOpcode(opCodeNum, numArgs int, flags string, icfunc IntcodeFunc) {
	ic.opcodes[opCodeNum] = intcodeopcode{args: numArgs, argflags: flags, icfunc: icfunc}
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
	for i := 1; i <= opcode.args; i++ {
		flag := ocflags % 10
		ocflags = ocflags / 10
		arg := ic.mem[ic.pc+i]
		if flag == 0 && opcode.argflags[i-1] == 'r' {
			arg = ic.mem[arg]
		}
		args[i-1] = arg
	}
	done, err = opcode.icfunc(ic, args)
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

// RunProgramChannelMode runs an intcode program in channel mode
func (ic *Intcode) RunProgramChannelMode(in, out chan int, err chan error, done chan bool) {
	ic.UseChannels = true
	ic.inchan = in
	ic.outchan = out
	go func() {
		var instrdone bool
		var instrerr error
		for !instrdone && instrerr == nil {
			instrdone, instrerr = ic.RunInstruction()
		}
		if instrerr != nil {
			err <- instrerr
		}
		done <- instrdone
	}()
}
