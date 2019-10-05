#!env python

import aoc
import collections
import pprint

# hex grid help at https://www.redblobgames.com/grids/hexagons/#coordinates

Hex = collections.namedtuple('Hex', ('x', 'y', 'z'))

AxialHex = collections.namedtuple('AxialHex', ('q', 'r'))

MOVES = {
    'n': Hex(x=0, y=1, z=-1),
    'ne': Hex(x=-1, y=1, z=0),
    'se': Hex(x=-1, y=0, z=1),
    's': Hex(x=0, y=-1, z=1),
    'sw': Hex(x=1, y=-1, z=0),
    'nw': Hex(x=1, y=0, z=-1),
}

AXIALMOVES = {
    'n': AxialHex(q=0, r=-1),
    's': AxialHex(q=0, r=1),
    'ne': AxialHex(q=-1, r=0),
    'sw': AxialHex(q=1, r=0),
    'se': AxialHex(q=-1, r=1),
    'nw': AxialHex(q=-1, r=1),
}


def parse_line(line):
    return line.split(',')


def parse_data(lines):
    return [parse_line(line) for line in lines][0]


def solve1(steps):
    pos = Hex(x=0, y=0, z=0)
    for step in steps:
        move = MOVES[step]
        pos = Hex(x=pos.x + move.x,
                  y=pos.y + move.y,
                  z=pos.z + move.z)
    return max(abs(value) for value in pos)


def solve2(steps):
    maxsteps = 0
    pos = Hex(x=0, y=0, z=0)
    for step in steps:
        move = MOVES[step]
        pos = Hex(x=pos.x + move.x,
                  y=pos.y + move.y,
                  z=pos.z + move.z)
        steps = max(abs(value) for value in pos)
        if steps > maxsteps:
            maxsteps = steps
    return maxsteps


def solve1and2axial(steps):
    maxsteps = 0
    pos = AxialHex(0, 0)
    for step in steps:
        move = AXIALMOVES[step]
        pos = AxialHex(q=pos.q + move.q, r=pos.r + move.r)
        steps = abs(sum(pos))
        if steps > maxsteps:
            maxsteps = steps
    return steps, maxsteps


lines = [
    'se,se,se,sw,sw'
]


if __name__ == '__main__':
    lines = aoc.input_lines(day=11)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
    pprint.pprint(solve1and2axial(data))
