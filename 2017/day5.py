#!env python

import aoc


def solve1(rows):
    steps = 0
    offset = 0
    jumps = list(rows)
    while 1:
        try:
            oldjump = jumps[offset]
        except IndexError:
            return steps
        steps += 1
        jumps[offset] += 1
        offset += oldjump

def solve2(rows):
    steps = 0
    offset = 0
    jumps = list(rows)
    while 1:
        try:
            oldjump = jumps[offset]
        except IndexError:
            return steps
        steps += 1
        if oldjump >= 3:
            jumps[offset] -= 1
        else:
            jumps[offset] += 1
        offset += oldjump


if __name__ == '__main__':
    rows = tuple([0, 3, 0, 1, -3])
    rows = tuple([int(row) for row in aoc.input_lines(day=5)])
    print(solve1(rows))
    print(solve2(rows))
