#!/bin/env python3


import hashlib
import re
import pprint


class Calculator(object):
    def __init__(self, salt):
        self.salt = salt
        self.hashes = {}

    def get_hash(self, index):
        if index not in self.hashes:
            hasher = hashlib.md5()
            hasher.update((self.salt + str(index)).encode('utf-8'))
            self.hashes[index] = hasher.hexdigest()
        return self.hashes[index]


def solve1(data):
    triple_regex = re.compile(r'(.)\1\1')
    calc = Calculator(data)
    return [(i, calc.get_hash(i), triple_regex.search(calc.get_hash(i))) for i in range(20)]


def solve2(data):
    pass


if __name__ == '__main__':

    data = 'zpqevtbw'
    data = 'abc'
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
