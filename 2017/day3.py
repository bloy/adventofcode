#!env python

import itertools
import collections


def rings():
    prevringmax = 0
    for ring in itertools.count(0):
        ringmax = ((ring*2)+1)**2
        ringmin = prevringmax + 1
        prevringmax = ringmax
        yield(ring, range(ringmin, ringmax+1))


def find_ring(num):
    for ring in rings():
        if num in ring[1]:
            return ring


def memorypositions():
    for ring, ringrange in rings():
        if ring == 0:
            yield (0, 0)
        xpos = itertools.chain(
            itertools.repeat(ring, ring*2),
            range(ring-1, -ring-1, -1),
            itertools.repeat(-ring, ring*2),
            range(-ring+1, ring+1)
        )
        ypos = itertools.chain(
            range(-ring+1, ring+1),
            itertools.repeat(ring, ring*2),
            range(ring-1, -ring-1, -1),
            itertools.repeat(-ring, ring*2)
        )
        for pos in zip(xpos, ypos):
            yield tuple(pos)


def solve1(num):
    ring, ringrange = find_ring(num)
    if ring == 0:
        return 0;
    rangedistance = zip(ringrange, itertools.cycle(
        itertools.chain(range(ring*2-1, ring-1, -1), range(ring+1, ring*2+1))))
    for testnum, dist in rangedistance:
        if testnum == num:
            return dist


def solve1alt(num):
    for i, pos in enumerate(memorypositions()):
        if i + 1 == num:
            return abs(pos[0]) + abs(pos[1])


def solve2(num):
    sums = collections.defaultdict(int)
    for i, pos in enumerate(memorypositions()):
        if i == 0:
            sums[pos] = 1;
        sums[pos] = sum(sums[(x, y)]
                        for x in range(pos[0]-1, pos[0]+2)
                        for y in range(pos[1]-1, pos[1]+2))
        if sums[pos] > num:
            return sums[pos];


if __name__ == '__main__':
    space = 361527
    examples = [space]
    for ex in examples:
        print(ex, solve1alt(ex))
        print(ex, solve2(ex))
