#!env python

import aoc
import collections
import pprint


def parse_data(lines):
    parts = collections.defaultdict(list)
    partlist = []
    for line in lines:
        ends = tuple(int(end) for end in line.split('/'))
        for num in ends:
            parts[num].append(ends)
        partlist.append(ends)
    return parts, partlist


def othernum(part, num):
    if part[0] == num:
        return part[1]
    return part[0]


def pathstr(path):
    return sum(sum(parts) for parts in path)


def solve1(data):
    parts, _ = data
    queue = collections.deque()
    queue.append((tuple(), 0))  # empty path with 'next num' == 0
    maxstr = 0
    while len(queue) > 0:
        path, num = queue.pop()
        strength = sum(sum(parts) for parts in path)
        if maxstr <= strength:
            maxstr = strength
        for next in filter(lambda x: x not in path, parts[num]):
            queue.append((path + (next,), othernum(next, num)))
    return maxstr


def solve2(data):
    parts, _ = data
    queue = collections.deque()
    queue.append((tuple(), 0))  # empty path with 'next num' == 0
    lengths = collections.defaultdict(list)
    while len(queue) > 0:
        path, num = queue.popleft()
        length = len(path)
        lengths[length].append(path)
        for next in filter(lambda x: x not in path, parts[num]):
            queue.append((path + (next,), othernum(next, num)))
    maxlen = max(lengths.keys())
    maxstr = 0
    for path in lengths[maxlen]:
        strength = sum(sum(parts) for parts in path)
        if strength > maxstr:
            maxstr = strength
    return maxstr


lines = [
    '0/2',
    '2/2',
    '2/3',
    '3/4',
    '3/5',
    '0/1',
    '10/1',
    '9/10',
]


if __name__ == '__main__':
    lines = aoc.input_lines(day=24)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
