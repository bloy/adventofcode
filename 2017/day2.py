#!env python

import itertools

import aoc


def checksum(row):
    return max(row) - min(row)


def solve1(rows):
    return sum(checksum(row) for row in rows)


def checksum2(row):
    for a, b in itertools.combinations(row, 2):
        if a % b == 0:
            return a // b
        if b % a == 0:
            return b // a;


def solve2(rows):
    return sum(checksum2(row) for row in rows)

if __name__ == '__main__':
    rows = [[int(col) for col in row.split('\t')]
            for row in aoc.input_lines(day=2)]
    print(solve1(rows))
    print(solve2(rows))
