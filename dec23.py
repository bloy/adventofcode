#!/usr/bin/env python
import pprint


def simple_machine(lines):
    registers = {
        'a': 1,
        'b': 0,
    }
    program_counter = 0
    while program_counter < len(lines):
        instruction, args = lines[program_counter].split(' ', 1)
        args = args.split(', ')
        pprint.pprint((program_counter, instruction, args))
        if instruction == 'hlf':
            register = args[0]
            registers[register] = registers[register] // 2
            program_counter += 1
        elif instruction == 'tpl':
            register = args[0]
            registers[register] = registers[register] * 3
            program_counter += 1
        elif instruction == 'inc':
            register = args[0]
            registers[register] = registers[register] + 1
            program_counter += 1
        elif instruction == 'jmp':
            offset = int(args[0])
            program_counter += offset
        elif instruction == 'jie':
            register = args[0]
            offset = int(args[1])
            if registers[register] % 2 == 0:
                program_counter += offset
            else:
                program_counter += 1
        elif instruction == 'jio':
            register = args[0]
            offset = int(args[1])
            if registers[register] == 1:
                program_counter += offset
            else:
                program_counter += 1
    return registers


if __name__ == '__main__':
    lines = list()
    with open('input/day_23') as f:
        for line in f:
            lines.append(line.strip())

    pprint.pprint(simple_machine(lines))
