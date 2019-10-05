#!env python

import aoc
import collections
import pprint


def parse_move(move):
    if move[0] == 's':
        return ('s', int(move[1:]))
    if move[0] == 'x':
        parts = move[1:].split('/')
        return ('x', int(parts[0]), int(parts[1]))
    if move[0] == 'p':
        parts = move[1:].split('/')
        return ('p', parts[0], parts[1])


def parse_data(lines):
    line = "".join(lines)
    return tuple([parse_move(move) for move in line.split(',')])


def do_move(dancers, move):
    if move[0] == 's':
        rotate = move[1]
        return dancers[-rotate:] + dancers[:-rotate]
    if move[0] == 'x':
        a = move[1]
        b = move[2]
    elif move[0] == 'p':
        a = dancers.index(move[1])
        b = dancers.index(move[2])
    dancers = list(dancers)
    dancers[a], dancers[b] = dancers[b], dancers[a]
    return "".join(dancers)


def solve1(moves):
    dancers = 'abcdefghijklmnop'
    for move in moves:
        dancers = do_move(dancers, move)
    return dancers


def solve2(moves):
    dancers = 'abcdefghijklmnop'
    seen = set()
    results = []
    for i in range(1000000000):
        results.append(dancers)
        if dancers in seen:
            print(dancers)
            previous = results.index(dancers)
            looplen = i - previous
            return results[(1000000000 - previous) % looplen] 
        seen.add(dancers)
        for move in moves:
            dancers = do_move(dancers, move)
    return dancers


lines = [
    's1,x3/4,pe/b'
]



if __name__ == '__main__':
    lines = aoc.input_lines(day=16)
    moves = parse_data(lines)
    pprint.pprint(solve1(moves))
    pprint.pprint(solve2(moves))
