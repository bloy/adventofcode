#!env python

import collections
import pprint

import aoc

Instruction = collections.namedtuple("Instruction",
                                     ('target', 'operation', 'delta',
                                      'compreg', 'compop', 'compnum'))


def parse_line(line):
    parts = line.split()
    return Instruction(
        target=parts[0],
        operation=parts[1],
        delta=int(parts[2]),
        compreg=parts[4],
        compop=parts[5],
        compnum=int(parts[6]))

def compare(value1, operation, value2):
    if operation == '>':
        return value1 > value2
    if operation == '<':
        return value1 < value2
    if operation == '>=':
        return value1 >= value2
    if operation == '<=':
        return value1 <= value2
    if operation == '==':
        return value1 == value2
    if operation == '!=':
        return value1 != value2
    raise ValueError("operation", operation)


def solve1(instructions):
    maxever = 0
    registers = collections.defaultdict(int)
    for instr in instructions:
        if compare(registers[instr.compreg], instr.compop, instr.compnum):
            multi = 1 if instr.operation == 'inc' else -1
            registers[instr.target] += (multi * instr.delta)
        currentmax = max(registers.values())
        if currentmax > maxever:
            maxever = currentmax
    return currentmax, maxever


if __name__ == '__main__':
    lines = (
        "b inc 5 if a > 1",
        "a inc 1 if b < 5",
        "c dec -10 if a >= 1",
        "c inc -20 if c == 10",
    )
    lines = aoc.input_lines(day=8)
    instructions = [parse_line(line) for line in lines]
    pprint.pprint(solve1(instructions))
