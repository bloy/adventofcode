#!/usr/bin/env python
import re
import pprint


def parse_sue_input(lines):
    sues = []
    for line in lines:
        line = line.strip()
        (sue_number, facts) = line.split(': ', 1)
        (_, sue_number) = sue_number.split(' ')
        facts = dict([fact.split(': ') for fact in facts.split(', ')])
        facts['number'] = sue_number
        for key in facts:
            facts[key] = int(facts[key])
        sues.append(facts)
    return sues


def deduce_part1(sues, clues):
    for clue in clues:
        sues = [sue for sue in sues
                if ((not sue.has_key(clue)) or
                    sue[clue] == clues[clue])]
    return sues

def valid_part2_sue(sue, clue, value):
    if not sue.has_key(clue):
        return True
    if clue in ('cats', 'trees'):
        if sue[clue] > value:
            return True
    elif clue in ('pomeranians', 'goldfish'):
        if sue[clue] < value:
            return True
    elif sue[clue] == value:
        return True
    else:
        return False


def deduce_part2(sues, clues):
    for clue in clues:
        sues = [sue for sue in sues if valid_part2_sue(sue, clue, clues[clue])]
    return sues

if __name__ == '__main__':
    with open('input/day_16') as lines:
        sues = parse_sue_input(lines)
    clues = {
        "children": 3,
        "cats": 7,
        "samoyeds": 2,
        "pomeranians": 3,
        "akitas": 0,
        "vizslas": 0,
        "goldfish": 5,
        "trees": 3,
        "cars": 2,
        "perfumes": 1
    }

    pprint.pprint(deduce_part1(sues, clues))
    pprint.pprint(deduce_part2(sues, clues))
