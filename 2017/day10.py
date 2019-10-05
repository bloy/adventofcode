#!env python

import functools
import pprint


def solve1(startarray, lengths):
    data = list(startarray)
    pos = 0
    skip = 0
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
    return data[0] * data[1]


def solve2(inputstr):
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



if __name__ == '__main__':
    # startarray = tuple(range(256))
    # lengths = (63, 144, 180, 149, 1, 255, 167, 84, 125, 65, 188, 0, 2, 254, 229, 24)
    # pprint.pprint(solve1(startarray, lengths))

    inputstr = "63,144,180,149,1,255,167,84,125,65,188,0,2,254,229,24"
    pprint.pprint(solve2(inputstr))
