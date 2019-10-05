#!env python

import collections
import functools
import pprint


def parse_data(lines):
    return lines


def knothash(inputstr):
    data = list(range(256))
    suffix = tuple([17, 31, 73, 47, 23])
    lengths = tuple((ord(c) for c in inputstr)) + suffix
    pos = 0
    skip = 0
    for iteration in range(64):
        for length in lengths:
            if pos + length >= len(data):
                endlen = len(data[pos:])
                endpos = pos + length - len(data)
                subdata = list(reversed(data[pos:] + data[:endpos]))
                data[pos:] = subdata[:endlen]
                data[:endpos] = subdata[endlen:]
            else:
                data[pos:pos + length] = list(reversed(data[pos:pos + length]))
            pos = (pos + length + skip) % len(data)
            skip += 1
    dense = []
    for rangestart in range(0, 255, 16):
        rangeend = rangestart + 16
        datarange = data[rangestart:rangeend]
        dense.append(functools.reduce(lambda x, y: x ^ y, datarange))
    dense = "".join("{:02x}".format(num) for num in dense)
    return dense


def generate_hashes(inputstr):
    return [knothash("{0}-{1}".format(inputstr, i)) for i in range(128)]


def solve1(data):
    hashes = data
    bits = 0
    for row in hashes:
        binrow = "{:04b}".format(int(row, base=16))
        bits += sum(int(c) for c in binrow)
    return bits


def get_group(bits, x, y):
    group_sites = set()
    seen = set()
    stack = collections.deque()
    stack.append((x, y))
    while stack:
        pos = stack.pop()
        if pos not in seen:
            seen.add(pos)
            if bits[pos[0]][pos[1]]:
                group_sites.add(pos)
                if pos[0] > 0:
                    stack.append((pos[0] - 1, pos[1]))
                if pos[1] > 0:
                    stack.append((pos[0], pos[1] - 1))
                if pos[0] < 127:
                    stack.append((pos[0] + 1, pos[1]))
                if pos[1] < 127:
                    stack.append((pos[0], pos[1] + 1))
    return group_sites


def solve2(data):
    bits = []
    for row in data:
        num = int(row, base=16)
        bitsrow = [num & (1 << i) == (1 << i) for i in range(128)]
        bits.append(bitsrow)
    seen = set()
    group_count = 0
    for x in range(128):
        for y in range(128):
            if bits[x][y] and (x, y) not in seen:
                group = get_group(bits, x, y)
                seen = seen.union(group)
                group_count += 1
    return group_count





if __name__ == '__main__':
    data = 'jzgqcdpd'
    hashes = generate_hashes(data)
    pprint.pprint(solve1(hashes))
    pprint.pprint(solve2(hashes))
