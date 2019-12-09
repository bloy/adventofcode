package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleIntcodeDay2Sample1(t *testing.T) {
	program := "1,9,10,3,2,3,11,0,99,30,40,50"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	ic.RunProgram(nil)
	expected := map[int64]int64{
		0: 3500, 1: 9, 2: 10,
		3: 70, 4: 2, 5: 3, 6: 11,
		7: 0, 8: 99, 9: 30, 10: 40,
		11: 50}
	assert.Equal(t, ic.mem, expected)
}

func TestSampleIntcodeDay2Input(t *testing.T) {
	program := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,9,19,23,1,23,5,27,2,27,10,31,1,6,31,35,1,6,35,39,2,9,39,43,1,6,43,47,1,47,5,51,1,51,13,55,1,55,13,59,1,59,5,63,2,63,6,67,1,5,67,71,1,71,13,75,1,10,75,79,2,79,6,83,2,9,83,87,1,5,87,91,1,91,5,95,2,9,95,99,1,6,99,103,1,9,103,107,2,9,107,111,1,111,6,115,2,9,115,119,1,119,6,123,1,123,9,127,2,127,13,131,1,131,9,135,1,10,135,139,2,139,10,143,1,143,5,147,2,147,6,151,1,151,5,155,1,2,155,159,1,6,159,0,99,2,0,14,0"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	ic.mem[1] = 12
	ic.mem[2] = 2
	ic.RunProgram(nil)
	var expected int64 = 6627023
	assert.Equal(t, ic.mem[0], expected)
}

func TestIntcodeDay5Sample1NoInput(t *testing.T) {
	program := "3,0,4,0,99"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	_, err = ic.RunProgram(nil)
	assert.NotNil(t, err)
}

func TestIntcodeDay5Sample1(t *testing.T) {
	program := "3,0,4,0,99"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	outputs, err := ic.RunProgram([]int64{12345})
	assert.Nil(t, err)
	assert.Equal(t, outputs, []int64{12345})
}

func TestIntcodeDay5Sample2(t *testing.T) {
	program := "1002,4,3,4,33"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	outputs, err := ic.RunProgram(nil)
	assert.Nil(t, err)
	assert.Equal(t, outputs, []int64{})
	assert.Equal(t, ic.mem, map[int64]int64{0: 1002, 1: 4, 2: 3, 3: 4, 4: 99})
}

func TestIntcodeDay5Sample3(t *testing.T) {
	program := "1101,100,-1,4,0"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	outputs, err := ic.RunProgram(nil)
	assert.Nil(t, err)
	assert.Equal(t, outputs, []int64{})
	assert.Equal(t, ic.mem, map[int64]int64{0: 1101, 1: 100, 2: -1, 3: 4, 4: 99})
}

func TestIntcodeDay9Sample1(t *testing.T) {
	program := "1101,100,23,1985,109,2000,109,19,204,-34,99"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	outputs, err := ic.RunProgram([]int64{})
	assert.Nil(t, err)
	assert.Equal(t, []int64{123}, outputs)
}

func TestIntcodeDay9Sample2(t *testing.T) {
	program := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	outputs, err := ic.RunProgram([]int64{})
	assert.Nil(t, err)
	expected := []int64{
		109, 1,
		204, -1,
		1001, 100, 1, 100,
		1008, 100, 16,
		101, 1006,
		101, 0,
		99}
	assert.Equal(t, outputs, expected)
}

func TestIntcodeDay9Sample3(t *testing.T) {
	program := "1102,34915192,34915192,7,4,7,99,0"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	outputs, err := ic.RunProgram([]int64{})
	assert.Nil(t, err)
	var expected int64 = 1219070632396864
	assert.Equal(t, outputs, []int64{expected})
}

func TestIntcodeDay9Sample4(t *testing.T) {
	program := "104,1125899906842624,99"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	outputs, err := ic.RunProgram([]int64{})
	assert.Nil(t, err)
	var expected int64 = 1125899906842624
	assert.Equal(t, outputs, []int64{expected})
}
