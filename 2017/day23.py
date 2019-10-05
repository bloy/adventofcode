#!env python

import aoc
import collections
import re
import pprint

INTRE = re.compile(r'^(-|\+)?[0-9]+$')


Point = collections.namedtuple('Point', ('x', 'y'))


def parse_data(lines):
    return [line.split() for line in lines]


def value_of(param, registers):
    if INTRE.match(param):
        return int(param)
    return registers[param]


def solve1(data):
    registers = collections.defaultdict(lambda: 0)
    pc = 0
    mulcount = 0
    while 0 <= pc < len(data):
        instr = data[pc][0]
        param1 = data[pc][1]
        param2 = data[pc][2]
        if instr == 'set':
            registers[param1] = value_of(param2, registers)
        elif instr == 'sub':
            registers[param1] -= value_of(param2, registers)
        elif instr == 'mul':
            registers[param1] *= value_of(param2, registers)
            mulcount += 1
        elif instr == 'jnz':
            if value_of(param1, registers) != 0:
                pc += value_of(param2, registers) - 1
        pc += 1
    return mulcount


def solve2():
    h = 0
    b = 84
    c = b
    b *= 100
    b += 100000
    c = b + 17000
    for e in range(b, c+1, 17):
        for d in range(2,e):
            if e % d == 0:
                h+= 1
                break
        print("b:", b, "c:", c, "d:", d,  "h:", h)
    return h


if __name__ == '__main__':
    lines = aoc.input_lines(day=23)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2())
