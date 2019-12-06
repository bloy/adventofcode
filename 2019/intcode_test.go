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
	expected := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
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
	expected := 6627023
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
	outputs, err := ic.RunProgram([]int{12345})
	assert.Nil(t, err)
	assert.Equal(t, outputs, []int{12345})
}

func TestIntcodeDay5Sample2(t *testing.T) {
	program := "1002,4,3,4,33"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	outputs, err := ic.RunProgram(nil)
	assert.Nil(t, err)
	assert.Equal(t, outputs, []int{})
	assert.Equal(t, ic.mem, []int{1002, 4, 3, 4, 99})
}

func TestIntcodeDay5Sample3(t *testing.T) {
	program := "1101,100,-1,4,0"
	ic, err := NewIntcodeFromInput(program)
	assert.Nil(t, err)
	ic.AddStandardOpcodes()
	outputs, err := ic.RunProgram(nil)
	assert.Nil(t, err)
	assert.Equal(t, outputs, []int{})
	assert.Equal(t, ic.mem, []int{1101, 100, -1, 4, 99})
}
