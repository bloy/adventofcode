#!/bin/env python3

import pprint


def solve(seed, size):
    a = seed
    while len(a) < size:
        b = "".join(reversed(a))
        b = b.replace('1', 'a').replace('0', '1').replace('a', '0')
        a = "{0}0{1}".format(a, b)

    checksum = a[:size]
    while len(checksum) % 2 == 0:
        new = [('1' if checksum[i] == checksum[i+1] else '0')
               for i in range(0, len(checksum), 2)]
        checksum = "".join(new)
    return checksum


if __name__ == '__main__':
    seed = '10011111011011001'
    size = 272
    #seed = '10000'
    #size = 20
    pprint.pprint(solve(seed, size))


    size = 35651584
    pprint.pprint(solve(seed, size))
