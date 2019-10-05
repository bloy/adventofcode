#!/usr/bin/env python3

import pprint
import re


def solve1(data):
    good_sequence_regexp = re.compile(r'(\w)(?!\1)(\w)\2\1')
    bad_sequence_regexp = re.compile(r'\[\w*(\w)(?!\1)(\w)\2\1\w*\]')
    return len(list(d for d in data
                    if good_sequence_regexp.search(d) and
                    not bad_sequence_regexp.search(d)))


def has_ssl(ip):
    for out_of_bracket in re.split('\[\w*\]', ip):
        for aba in re.finditer(r'(?=((\w)(?!\2)\w\2))', out_of_bracket):
            aba = aba.group(1)
            bab = '{0}{1}{2}'.format(aba[1], aba[0], aba[1])
            if any((hypernet.group(0).find(bab) > -1)
                   for hypernet in re.finditer(r'\[\w*\]', ip)):
                return True
    return False


def solve2(data):
    return len(list(d for d in data if has_ssl(d)))

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
