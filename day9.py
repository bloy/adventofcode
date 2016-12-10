#!/usr/bin/env python3

import pprint
import re


COMPRESSION_REGEX = r'\((\d+)x(\d+)\)'


def solve1(data):
    decompressed = ""
    split = re.split(COMPRESSION_REGEX, data, maxsplit=1)
    while split[0] != data:
        decompressed += split[0]
        (run, count, data) = (int(split[1]), int(split[2]), split[3])
        decompressed += (data[:run] * count)
        data = data[run:]
        split = re.split(COMPRESSION_REGEX, data, maxsplit=1)
    decompressed += split[0]
    return len(decompressed)


def solve2(data):
    return 0


if __name__ == '__main__':
    with open('day9_input.txt') as f:
        data = f.read().strip()

    # data = 'X(8x2)(3x3)AB(2x2)CY'

    print(solve1(data))
    print(solve2(data))
