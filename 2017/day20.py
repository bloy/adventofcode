#!env python

import aoc
import collections
import itertools
import pprint


Point = collections.namedtuple('Point', ['x', 'y', 'z'])


def part2point(part):
    x, y, z = part[3:-1].split(',')
    return Point(x=int(x), y=int(y), z=int(z))


def parse_data(lines):
    particles = []
    for num, line in enumerate(lines):
        p, v, a = line.split(', ')
        particles.append({
            "num": num,
            "p": part2point(p),
            "v": part2point(v),
            "a": part2point(a),
        })
    return particles


def num_at_tick(tick, p, v, a):
    delta = sum(v + t * a for t in range(1, tick+1))
    return p + delta


def pos_at_tick(tick, particle):
    return Point(
        x=num_at_tick(tick, particle['p'].x, particle['v'].x, particle['a'].x),
        y=num_at_tick(tick, particle['p'].y, particle['v'].y, particle['a'].y),
        z=num_at_tick(tick, particle['p'].z, particle['v'].z, particle['a'].z),
    )


def distance_at_tick(tick, particle):
    pos = pos_at_tick(tick, particle)
    return abs(pos.x) + abs(pos.y) + abs(pos.z)


def solve1(data):
    result = []
    for particle in data:
        result.append((
            particle['num'], distance_at_tick(1000, particle)))
    return min(result, key=lambda v: v[1])


def solve2(data):
    particles = data
    count = len(particles)
    deltacount = 0
    for t in itertools.count():
        positions = [pos_at_tick(t, p) for p in particles]
        for p in set(positions):
            if positions.count(p) > 1:
                while p in positions:
                    i = positions.index(p)
                    del positions[i]
                    del particles[i]
        if count != len(particles):
            count = len(particles)
            deltacount = 0
        else:
            deltacount += 1
        print("time:", t, "time since change:", deltacount,
              "num particles:", len(particles))
        if deltacount >= 1000:
            return len(particles)


lines = [
    'p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>',
    'p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>'
]


if __name__ == '__main__':
    lines = aoc.input_lines(day=20)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
