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


def decompressed_length(compressed):
    decom_len = 0
    split = re.split(COMPRESSION_REGEX, compressed, maxsplit=1)
    while split[0] != compressed:
        decom_len += len(split[0])
        (run, count, compressed) = (int(split[1]), int(split[2]), split[3])
        decom_len += (decompressed_length(compressed[:run]) * count)
        compressed = compressed[run:]
        split = re.split(COMPRESSION_REGEX, compressed, maxsplit=1)
    decom_len += len(split[0])
    return decom_len


def solve2(data):
    return decompressed_length(data)


if __name__ == '__main__':
    with open('day9_input.txt') as f:
        data = f.read().strip()

    # data = '(27x12)(20x12)(13x14)(7x10)(1x12)A'

    print(solve1(data))
    print(solve2(data))
