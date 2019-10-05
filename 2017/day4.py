#!env python

import aoc

import itertools


def solve1(rows):
    count = 0
    for row in rows:
        words = row.split(' ')
        if len(words) == len(set(words)):
            count += 1
    return count


def valid2(phrase):
    wordset = set()
    for word in phrase.split(' '):
        if word in wordset:
            return False
        for scramble in itertools.permutations(word):
            wordset.add("".join(scramble))
    return True


def solve2(rows):
    count = 0
    for row in rows:
        if valid2(row):
            count += 1
    return count


def solve2alt(rows):
    count = 0
    for row in rows:
        words = ["".join(sorted(word)) for word in row.split(' ')]
        if len(words) == len(set(words)):
            count += 1
    return count


if __name__ == '__main__':
    rows = [row for row in aoc.input_lines(day=4)]
    print(solve1(rows))
    print(solve2alt(rows))
