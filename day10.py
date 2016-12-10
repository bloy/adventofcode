#!/usr/bin/env python3

import collections
import pprint
import re


VALUE_REGEX = re.compile('^value (\d+) goes to (output|bot) (\d+)$')
INSTR_REGEX = re.compile('^bot (\d+) gives low to (output|bot) (\d+) and high to (output|bot) (\d+)$')


def solve1(data):
    bots = collections.defaultdict(set)
    outputs = collections.defaultdict(set)
    instructions = collections.defaultdict(None)

    for line in data:
        match = VALUE_REGEX.match(line)
        if match:
            value = int(match.group(1))
            if match.group(2) == 'bot':
                bots[match.group(3)].add(value)
        match = INSTR_REGEX.match(line)
        if match:
            instructions[match.group(1)] = {
                'low': (match.group(2), match.group(3)),
                'high': (match.group(4), match.group(5)),
            }

    while True:
        target_bots = [bot for bot in bots if bots[bot] == set([61, 17])]
        if target_bots:
            return target_bots
        for bot in [b for b in bots if len(bots[b]) == 2]:
            values = {'high': max(bots[bot]),
                      'low': min(bots[bot])}
            bots[bot] == set()
            instruction = instructions[bot]

            for sort in values:
                if instruction[sort][0] == 'bot':
                    bots[instruction[sort][1]].add(values[sort])
                else:
                    outputs[instruction[sort][1]].add(values[sort])


def solve2(data):
    bots = collections.defaultdict(set)
    outputs = collections.defaultdict(set)
    instructions = collections.defaultdict(None)

    for line in data:
        match = VALUE_REGEX.match(line)
        if match:
            value = int(match.group(1))
            if match.group(2) == 'bot':
                bots[match.group(3)].add(value)
        match = INSTR_REGEX.match(line)
        if match:
            instructions[match.group(1)] = {
                'low': (match.group(2), match.group(3)),
                'high': (match.group(4), match.group(5)),
            }

    while True:
        if '0' in outputs and '1' in outputs and '2' in outputs:
            break;
        multi_input_bots = [bot for bot in bots if len(bots[bot]) == 2]
        for bot in multi_input_bots:
            values = {'high': max(bots[bot]),
                      'low': min(bots[bot])}
            bots[bot] == set()
            instruction = instructions[bot]

            for sort in values:
                if instruction[sort][0] == 'bot':
                    bots[instruction[sort][1]].add(values[sort])
                else:
                    outputs[instruction[sort][1]].add(values[sort])

    return sum(outputs['1']) * sum(outputs['2']) * sum(outputs['0'])


if __name__ == '__main__':
    with open('day10_input.txt') as f:
        data = [line.strip() for line in f]

    # data = [
        # 'value 5 goes to bot 2',
        # 'bot 2 gives low to bot 1 and high to bot 0',
        # 'value 3 goes to bot 1',
        # 'bot 1 gives low to output 1 and high to bot 0',
        # 'bot 0 gives low to output 2 and high to output 0',
        # 'value 2 goes to bot 2',
    # ]

    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
