#!/usr/bin/env python3

import pprint

def solve1(data):
    pass

def solve2(data):
    pass

if __name__ == '__main__':
    with open('day7_input.txt') as f:
        data = [line for line in f.read().split("\n") if line != '']
    data = [
        "abba[mnop]qrst"  # yes
        "abcd[bddb]xyyx"  # no
        "aaaa[qwer]tyui"  # no
        "ioxxoj[asdfgh]zxcvbni" # yes
    ]

    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
