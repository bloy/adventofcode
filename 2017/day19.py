#!env python

import aoc
import collections
import pprint


def parse_data(lines):
    return [line for line in lines]


def print_map(data):
    print("\n".join(data))


def char_at(data, x, y, direction=None):
    if direction:
        if direction == 'S':
            y = y + 1
        elif direction == 'N':
            y = y - 1
        elif direction == 'W':
            x = x + 1
        elif direction == 'E':
            x = x - 1
    try:
        return data[y][x]
    except IndexError:
        return " "


def solve1(data):
    x = data[0].index('|')
    y = 0
    stepcount = 0
    direction = 'S'
    letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
    path = ""
    current = char_at(data, x, y)
    while current != " ":
        if current in letters:
            path = path + current
        if char_at(data, x, y, direction) == ' ':
            if direction in 'SN':
                if char_at(data, x, y, 'E') != ' ':
                    direction = 'E'
                else:
                    direction = 'W'
            else:
                if char_at(data, x, y, 'N') != ' ':
                    direction = 'N'
                else:
                    direction = 'S'
        if direction == 'S':
            y = y + 1
        elif direction == 'N':
            y = y - 1
        elif direction == 'W':
            x = x + 1
        elif direction == 'E':
            x = x - 1
        current = char_at(data, x, y)
        stepcount += 1
    return "".join(path), stepcount


def solve2(data):
    pass


lines = [
    "     |          ",
    "     |  +--+    ",
    "     A  |  C    ",
    " F---|----E|--+ ",
    "     |  |  |  D ",
    "     +B-+  +--+ ",
    "                ",
]


if __name__ == '__main__':
    lines = aoc.input_lines(day=19, strip=False)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
