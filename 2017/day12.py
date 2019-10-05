#!env python

import aoc
import collections
import pprint


def parse_data(lines):
    connections = collections.defaultdict(set)
    for line in lines:
        source, dests = line.split(' <-> ')
        source = int(source)
        for dest in dests.split(', '):
            dest = int(dest)
            connections[source].add(dest)
            connections[dest].add(source)
    return connections


def get_group(data, group_member):
    seen = set()
    stack = collections.deque()
    stack.append(group_member)
    while len(stack):
        item = stack.pop()
        seen.add(item)
        for connection in data[item]:
            if connection not in seen:
                stack.append(connection)
    return seen


def solve1(data):
    group = get_group(data, 0)
    return len(group)


def solve2(data):
    seen = set()
    groups = []
    for key in data.keys():
        if key not in seen:
            group = get_group(data, key)
            groups.append(group)
            seen = seen.union(group)
    return len(groups)


lines = [
    "0 <-> 2",
    "1 <-> 1",
    "2 <-> 0, 3, 4",
    "3 <-> 2, 4",
    "4 <-> 2, 3, 6",
    "5 <-> 6",
    "6 <-> 4, 5",
]


if __name__ == '__main__':
    lines = aoc.input_lines(day=12)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
