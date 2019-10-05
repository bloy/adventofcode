#!/usr/bin/env python3

import collections
import itertools
import re


class Instruction(collections.namedtuple('Instruction', 'keyword arg1 arg2')):
    __slots__ = ()
    INSTR_REGEX = re.compile(
        r'^\s*(?P<instr>tgl|out|cpy|inc|dec|jnz)\s+(?P<arg1>\S+)(?:\s+(?P<arg2>\S+))?\s*$')
    DIGITS_REGEX = re.compile(r'^([+-]?\d+)$')

    toggle_keywords = {
        'cpy': 'jnz',
        'jnz': 'cpy',
        'inc': 'dec',
        'dec': 'inc',
        'tgl': 'inc',
        'out': 'inc',
    }

    def toggle(self):
        return self.__class__(self.toggle_keywords[self.keyword],
                              self.arg1, self.arg2)

    @classmethod
    def parse(cls, line):
        match = cls.INSTR_REGEX.match(line)
        if match:
            groups = match.groupdict()
            if cls.DIGITS_REGEX.match(groups['arg1']):
                groups['arg1'] = int(groups['arg1'])
            if groups['arg2'] and cls.DIGITS_REGEX.match(groups['arg2']):
                groups['arg2'] = int(groups['arg2'])
            return cls(keyword=groups['instr'],
                       arg1=groups['arg1'], arg2=groups['arg2'])

    def __str__(self):
        return "{0:<4} {1:<3} {2:<3}".format(self.keyword,
                                            self.arg1,
                                            self.arg2 if self.arg2 else "")


class Machine(object):
    def __init__(self, instructions, a=0, b=0, c=0, d=0):
        self.instructions = [Instruction.parse(line) for line in instructions]
        self.output = []
        self.pc = 0
        self.registers = { 'a': a, 'b': b, 'c': c, 'd': d }
        self.instruction_count = 0
        self.keywords = {
            'cpy': self._cpy,
            'inc': self._inc,
            'dec': self._dec,
            'tgl': self._tgl,
            'jnz': self._jnz,
            'out': self._out,
        }

    def __str__(self):
        return "pc: {0:<2d} a: {1:<6d} b: {2:<6d} c: {3:<6d} d: {4:<6d}".format(
            self.pc,
            self.registers['a'],
            self.registers['b'],
            self.registers['c'],
            self.registers['d'])

    def next_instruction(self):
        while self.pc < len(self.instructions):
           yield self.instructions[self.pc]

    def run(self):
        for instr in self.next_instruction():
            self.keywords[instr.keyword](instr.arg1, instr.arg2)
            if len(self.output) == 10:
                break
        return self.registers

    def _cpy(self, arg1, arg2):
        if arg2 in self.registers:
            if isinstance(arg1, int):
                value = arg1
            else:
                value = self.registers[arg1]
            self.registers[arg2] = value
        self.pc += 1

    def _inc(self, arg1, arg2):
        if arg1 in self.registers:
           self.registers[arg1] += 1
        self.pc += 1

    def _dec(self, arg1, arg2):
        if arg1 in self.registers:
           self.registers[arg1] -= 1
        self.pc += 1

    def _tgl(self, arg1, arg2):
        value = self.registers[arg1] if arg1 in self.registers else arg1
        if self.pc + value >= 0 and self.pc + value < len(self.instructions):
            instr = self.instructions[self.pc + value]
            self.instructions[self.pc + value] = instr.toggle()
        self.pc += 1

    def _jnz(self, arg1, arg2):
        check = self.registers[arg1] if arg1 in self.registers else arg1
        value = self.registers[arg2] if arg2 in self.registers else arg2
        if check == 0:
            self.pc += 1
        else:
            self.pc += value

    def _out(self, arg1, arg2):
        value = self.registers[arg1] if arg1 in self.registers else arg1
        self.output.append(value)
        self.pc += 1


if __name__ == '__main__':
    data = [
        "cpy a d",
        "cpy 7 c",
        "cpy 362 b",
        "inc d",
        "dec b",
        "jnz b -2",
        "dec c",
        "jnz c -5",
        "cpy d a",
        "jnz 0 0",
        "cpy a b",
        "cpy 0 a",
        "cpy 2 c",
        "jnz b 2",
        "jnz 1 6",
        "dec b",
        "dec c",
        "jnz c -4",
        "inc a",
        "jnz 1 -7",
        "cpy 2 b",
        "jnz c 2",
        "jnz 1 4",
        "dec b",
        "dec c",
        "jnz 1 -4",
        "jnz 0 0",
        "out b",
        "jnz a -19",
        "jnz 1 -21",
    ]

    expected1 = list(itertools.islice(itertools.cycle((0, 1)), 10))

    for i in itertools.count():
        machine = Machine(data, a=i)
        machine.run()
        print(i, machine.output)
        if machine.output == expected1:
            print("i = {0}".format(i))
            break


