#!env python

import aoc
import collections
import pprint


Point = collections.namedtuple('Point', ('x', 'y'))

NORTH = Point(x=0, y=-1)
SOUTH = Point(x=0, y=1)
EAST = Point(x=-1, y=0)
WEST = Point(x=1, y=0)


DIRECTIONS = {
    (NORTH, True): WEST,
    (NORTH, False): EAST,
    (SOUTH, True): EAST,
    (SOUTH, False): WEST,
    (EAST, True): NORTH,
    (EAST, False): SOUTH,
    (WEST, True): SOUTH,
    (WEST, False): NORTH,
}


DIRECTIONS2 = {
    (NORTH, 'C'): EAST,
    (SOUTH, 'C'): WEST,
    (EAST, 'C'): SOUTH,
    (WEST, 'C'): NORTH,
    (NORTH, 'W'): NORTH,
    (SOUTH, 'W'): SOUTH,
    (EAST, 'W'): EAST,
    (WEST, 'W'): WEST,
    (NORTH, 'I'): WEST,
    (SOUTH, 'I'): EAST,
    (EAST, 'I'): NORTH,
    (WEST, 'I'): SOUTH,
    (NORTH, 'F'): SOUTH,
    (SOUTH, 'F'): NORTH,
    (EAST, 'F'): WEST,
    (WEST, 'F'): EAST,
}

NEXT_STATE2 = {
    'C': 'W',
    'W': 'I',
    'I': 'F',
    'F': 'C',
}


def parse_data(lines):
    grid = [line for line in lines]
    xsize = len(grid[0])
    ysize = len(grid)
    xmid = xsize // 2
    ymid = xsize // 2
    infected = list(Point(x=x-xmid, y=y-ymid)
                   for y in range(ysize)
                   for x in range(xsize)
                   if grid[y][x] == '#')
    return infected


def solve1(data):
    infected = set(data)
    infectioncount = 0
    pos = Point(0, 0)
    direction = NORTH
    for burst in range(10000):
        position_infected = pos in infected
        direction = DIRECTIONS[(direction, position_infected)]
        if position_infected:
            infected.remove(pos)
        else:
            infectioncount += 1
            infected.add(pos)
        pos = Point(pos.x + direction.x, pos.y + direction.y)
    return infectioncount


def solve2(data):
    infectioncount = 0
    grid = collections.defaultdict(lambda: 'C')
    for d in data:
        grid[d] = 'I'
    pos = Point(0, 0)
    direction = NORTH
    steps = 10000000
    for burst in range(steps):
        if burst % 100000 == 0:
            print(burst, (burst * 100 / steps), infectioncount)
        pos_status = grid[pos]
        direction = DIRECTIONS2[(direction, pos_status)]
        grid[pos] = NEXT_STATE2[pos_status]
        if grid[pos] == 'I':
            infectioncount += 1
        pos = Point(pos.x + direction.x, pos.y + direction.y)
    return infectioncount


lines = [
    '..#',
    '#..',
    '...',
]


if __name__ == '__main__':
    lines = aoc.input_lines(day=22)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
