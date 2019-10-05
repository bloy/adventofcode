#!/usr/bin/env python3

import pprint
import collections

def solve1(data):
    return "".join([collections.Counter(position).most_common()[0][0]
            for position in zip(*data)])


def solve2(data):
    return "".join([collections.Counter(position).most_common()[-1][0]
            for position in zip(*data)])


if __name__ == '__main__':
    with open('day6_input.txt') as f:
        data = [line for line in f.read().split("\n") if line != '']

    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
