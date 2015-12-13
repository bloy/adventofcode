#!/usr/bin/env python
import re
import itertools
import pprint


def parse_lines(lines):
    parser_re = re.compile(r'^(?P<name1>\w+) would '
                           r'(?P<sign>gain|lose) (?P<number>\d+) happiness '
                           r'units by sitting next to (?P<name2>\w+)\.$')
    matrix = dict()
    guests = set()
    for line in lines:
        line = line.strip()
        groups = parser_re.match(line).groupdict()
        number = int(groups['number'])
        if groups['sign'] == 'lose':
            number = -1 * number
        matrix[(groups['name1'], groups['name2'])] = number
        guests.add(groups['name1'])
        guests.add(groups['name2'])
    return guests, matrix


if __name__ == '__main__':
    in_lines = [
        'Alice would gain 54 happiness units by sitting next to Bob.',
        'Alice would lose 79 happiness units by sitting next to Carol.',
        'Alice would lose 2 happiness units by sitting next to David.',
        'Bob would gain 83 happiness units by sitting next to Alice.',
        'Bob would lose 7 happiness units by sitting next to Carol.',
        'Bob would lose 63 happiness units by sitting next to David.',
        'Carol would lose 62 happiness units by sitting next to Alice.',
        'Carol would gain 60 happiness units by sitting next to Bob.',
        'Carol would gain 55 happiness units by sitting next to David.',
        'David would gain 46 happiness units by sitting next to Alice.',
        'David would lose 7 happiness units by sitting next to Bob.',
        'David would gain 41 happiness units by sitting next to Carol.',
    ]
    guests, matrix = parse_lines(in_lines)
    pprint.pprint(guests)
    pprint.pprint(matrix)
