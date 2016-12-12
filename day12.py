#!/usr/bin/env python3

import re

INSTR_REGEX = re.compile(
    r'^\s*(?P<instr>cpy|inc|dec|jnz)\s+(?P<arg1>\S+)(?:\s+(?P<arg2>\S+))?\s*$')
DIGITS_REGEX = re.compile(r'^(\d+)$')


def parse_instruction(instr):
    match = INSTR_REGEX.match(instr)
    if match:
        parts = match.groupdict()
        if DIGITS_REGEX.match(parts['arg1']):
            parts['arg1'] = int(parts['arg1'])
        if parts['arg2'] and DIGITS_REGEX.match(parts['arg2']):
            parts['arg2'] = int(parts['arg2'])
        return parts
    else:
        return None


def solve1(instructions):
    registers = {
        'a': 0,
        'b': 0,
        'c': 0,
        'd': 0
    }
    pc = 0
    while pc < len(instructions):
        instr = parse_instruction(instructions[pc])
        arg1 = instr['arg1']
        arg2 = instr['arg2']
        keyword = instr['instr']
        print(pc, keyword, arg1, arg2, registers)
        if keyword == 'cpy':
            if isinstance(arg1, int):
                registers[arg2] = arg1
            else:
                registers[arg2] = registers[arg1]
            pc += 1
        elif keyword in ('inc', 'dec'):
            amount = 1 if keyword == 'inc' else -1
            registers[arg1] += amount
            pc += 1
        elif keyword == 'jnz':
            check = arg1 if isinstance(arg1, int) else registers[arg1]
            if check == 0:
                pc += 1
            else:
                pc += (int(arg2))
    return registers


def solve2(data):
    pass


if __name__ == '__main__':
    data = [
        "cpy 1 a",
        "cpy 1 b",
        "cpy 26 d",
        "jnz c 2",
        "jnz 1 5",
        "cpy 7 c",
        "inc d",
        "dec c",
        "jnz c -2",
        "cpy a c",
        "inc a",
        "dec b",
        "jnz b -2",
        "cpy c b",
        "dec d",
        "jnz d -6",
        "cpy 19 c",
        "cpy 14 d",
        "inc a",
        "dec d",
        "jnz d -2",
        "dec c",
        "jnz c -5",
    ]

    # data = [
        # "cpy 41 a",
        # "inc a",
        # "inc a",
        # "dec a",
        # "jnz a 2",
        # "dec a",
    # ]
    print(solve1(data))
    print(solve2(data))
