#!env python

import aoc
import collections
import pprint
import re


ISINT = re.compile(r'^-?[0-9]+$')


def parse_data(lines):
    return [line.split() for line in lines]


def valueof(v, registers):
    if v is None:
        return None
    if ISINT.match(v):
        return int(v)
    return registers[v]


def solve1(data):
    registers = collections.defaultdict(int)
    pc = 0
    soundplayed = 0
    while 0 <= pc < len(data):
        instr = data[pc][0]
        v1 = data[pc][1]
        v2 = data[pc][2] if len(data[pc]) > 2 else None
        pc += 1

        if instr == 'snd':
            soundplayed = valueof(v1, registers)
        elif instr == 'rcv':
            if valueof(v1, registers) != 0:
                return ('Last sound played', soundplayed)
        elif instr == 'set':
            registers[v1] = valueof(v2, registers)
        elif instr == 'add':
            registers[v1] += valueof(v2, registers)
        elif instr == 'mul':
            registers[v1] *= valueof(v2, registers)
        elif instr == 'mod':
            registers[v1] = registers[v1] % valueof(v2, registers)
        elif instr == 'jgz':
            if valueof(v1, registers) > 0:
                pc += valueof(v2, registers) - 1

    return "terminated"


def program(data, pid, rcvqueue, sndqueue):
    registers = collections.defaultdict()
    registers['p'] = pid
    pc = 0
    sendcount = 0
    terminated = False
    while 0 <= pc < len(data) and not terminated:
        instr = data[pc][0]
        v1 = data[pc][1]
        v2 = data[pc][2] if len(data[pc]) > 2 else None
        pc += 1

        if instr == 'snd':
            sndqueue.appendleft(valueof(v1, registers))
            sendcount += 1
        elif instr == 'rcv':
            if len(rcvqueue) == 0:
                yield sendcount
            try:
                registers[v1] = rcvqueue.pop()
            except IndexError:
                terminated = True
        elif instr == 'set':
            registers[v1] = valueof(v2, registers)
        elif instr == 'add':
            registers[v1] += valueof(v2, registers)
        elif instr == 'mul':
            registers[v1] *= valueof(v2, registers)
        elif instr == 'mod':
            registers[v1] = registers[v1] % valueof(v2, registers)
        elif instr == 'jgz':
            if valueof(v1, registers) > 0:
                pc += valueof(v2, registers) - 1
    yield sendcount


def solve2(data):
    queues = [collections.deque(), collections.deque()]
    programs = [program(data, 0, queues[0], queues[1]),
                program(data, 1, queues[1], queues[0])]
    current = 0
    returns = [None, None]
    while 1:
        try:
            returns[current] = next(programs[current])
        except StopIteration:
            return returns
        current = (current + 1) % 2



lines = [
    'set a 1',
    'add a 2',
    'mul a a',
    'mod a 5',
    'snd a',
    'set a 0',
    'rcv a',
    'jgz a -1',
    'set a 1',
    'jgz a -2',
]


if __name__ == '__main__':
    lines = aoc.input_lines(day=18)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
