#!/usr/bin/env python3

import pprint
import re

good_sequence_regexp = re.compile(r'(\w)(?!\1)(\w)\2\1')
bad_sequence_regexp = re.compile(r'\[\w*(\w)(?!\1)(\w)\2\1\w*\]')

part2_regex1 = re.compile(r'\[\w*(\w)(?!\1)(\w)\1\w*\]\w*\2\1\2')
part2_regex2 = re.compile(r'(\w)(?!\1)(\w)\1\w*\[\w*\2\1\2\w*\]')


def solve1(data):
    return len(list(d for d in data
                    if good_sequence_regexp.search(d) and
                    not bad_sequence_regexp.search(d)))


def solve2(data):
    return len(list(d for d in data if part2_regex1.search(d) or part2_regex2.search(d)))

if __name__ == '__main__':
    with open('day7_input.txt') as f:
        data = [line for line in f.read().split("\n") if line != '']

    # data = [
        # "aba[bab]xyz",   # yes
        # "xyx[xyx]xyx",   # no
        # "aaa[kek]eke",   # yes
        # "zazbz[bzb]cdb", # yes
    # ]

    # data = [
        # "abba[mnop]qrst",  # yes
        # "abcd[bddb]xyyx", # no
        # "aaaa[qwer]tyui",  # no
        # "ioxxoj[asdfgh]zxcvbni", # yes
    # ]

    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
