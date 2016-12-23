#!/usr/bin/env python3

import re
import collections



class Instruction(collections.namedtuple('Instruction', 'keyword arg1 arg2')):
    __slots__ = ()
    INSTR_REGEX = re.compile(
        r'^\s*(?P<instr>tgl|cpy|inc|dec|jnz)\s+(?P<arg1>\S+)(?:\s+(?P<arg2>\S+))?\s*$')
    DIGITS_REGEX = re.compile(r'^([+-]?\d+)$')

    toggle_keywords = {
        'cpy': 'jnz',
        'jnz': 'cpy',
        'inc': 'dec',
        'dec': 'inc',
        'tgl': 'inc',
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
        self.pc = 0
        self.registers = { 'a': a, 'b': b, 'c': c, 'd': d }
        self.instruction_count = 0
        self.keywords = {
            'cpy': self._cpy,
            'inc': self._inc,
            'dec': self._dec,
            'tgl': self._tgl,
            'jnz': self._jnz,
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
            if self.pc == 4 and (instr.keyword == 'cpy' and
                                 self.instructions[5].keyword == 'inc' and
                                 self.instructions[6].keyword == 'dec' and
                                 self.instructions[7].keyword == 'jnz' and
                                 self.instructions[8].keyword == 'dec' and
                                 self.instructions[9].keyword == 'jnz'):
                print("Multiplication hack", self)
                target = self.instructions[self.pc+1].arg1
                source1 = instr.arg1
                source2 = self.instructions[self.pc+4].arg1
                tmp = instr.arg2
                self.registers[target] = self.registers[source1] * self.registers[source2]
                self.registers[tmp] = 0
                self.registers[source2] = 0
                self.pc += 6
            else:
                print(instr, self)
                self.keywords[instr.keyword](instr.arg1, instr.arg2)
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


if __name__ == '__main__':
    data = [
        "cpy a b", # 0
        "dec b",   # 1
        "cpy a d", # 2
        "cpy 0 a", # 3
        "cpy b c", # 4 start of mul
        "inc a",   # 5
        "dec c",   # 6
        "jnz c -2", # 7
        "dec d",   # 8
        "jnz d -5",# 9 end of mul
        "dec b",   # 10
        "cpy b c",
        "cpy c d",
        "dec d",
        "inc c",
        "jnz d -2",
        "tgl c",
        "cpy -16 c",
        "jnz 1 c",
        "cpy 75 c",
        "jnz 72 d",
        "inc a",
        "inc d",
        "jnz d -2",
        "inc c",
        "jnz c -5",
    ]

    # data = [
        # "cpy 2 a",
        # "tgl a",
        # "tgl a",
        # "tgl a",
        # "cpy 1 a",
        # "dec a",
        # "dec a",
    # ]
    machine = Machine(data, a=7)
    print(machine.run())
    machine = Machine(data, a=12)
    print(machine.run())
