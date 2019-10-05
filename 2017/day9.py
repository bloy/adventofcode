#!env python

import aoc
import re
import pprint


def parse_data(lines):
    return [line for line in lines][0]


def solve1(data):
    cancelre = re.compile(r'!.')
    garbagere = re.compile(r'<.*?>')
    data = cancelre.sub('', data)
    data = garbagere.sub('', data)
    i = 0
    groupcount = 0
    score = 0
    while i < len(data):
        c = data[i]
        if c == '{':
            groupcount += 1
        elif c == '}':
            score += groupcount
            groupcount -= 1
        i += 1
    return score


def solve2(data):
    cancelre = re.compile(r'!.')
    garbagere = re.compile(r'<.*?>')
    data = cancelre.sub('', data)
    count = 0
    for match in garbagere.finditer(data):
        length = match.end() - match.start() - 2
        count += length
    return count



if __name__ == '__main__':
    lines = aoc.input_lines(day=9)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
