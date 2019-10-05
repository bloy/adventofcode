#!env python

import collections
import pprint


def parse_data(lines):
    return lines


def solve1(data):
    buf = [0]
    pos = 0
    for i in range(1, 2018):
        pos = (pos + data) % len(buf) + 1
        buf = buf[:pos] + [i] + buf[pos:]
    return buf[(buf.index(2017) + 1) % len(buf)]


def solve2(data):
    pos = 0
    after0 = 0
    for i in range(1,50000001):
        pos = (pos + data) % i + 1
        if pos == 1:
            after0 = i
    return after0

data = 367


if __name__ == '__main__':
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
