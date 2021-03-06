#!/bin/env python3


import hashlib
import re
import pprint


class Calculator(object):
    def __init__(self, salt):
        self.salt = salt
        self.hashes = {}

    def md5(self, value):
        hasher = hashlib.md5()
        hasher.update(str(value).encode('utf-8'))
        return hasher.hexdigest()

    def get_hash(self, index):
        if index not in self.hashes:
            md5 = self.md5(self.salt + str(index))
            self.hashes[index] = md5
        return self.hashes[index]

    def get_stretched_hash(self, index):
        if index not in self.hashes:
            md5 = self.md5(self.salt + str(index))
            for i in range(2016):
                md5 = self.md5(md5)
            self.hashes[index] = md5
        return self.hashes[index]


def solve1(data):
    triple_regex = re.compile(r'(.)\1\1')
    calc = Calculator(data)
    keys = []
    index = -1
    while len(keys) < 64:
        index += 1
        md5 = calc.get_hash(index)
        match = triple_regex.search(md5)
        if match:
            subindex = index + 1
            seq = match.group(1) * 5
            if any(seq in calc.get_hash(subindex + i) for i in range(1000)):
                print("found key {0} at index {1}".format(md5, index))
                keys.append(md5)
    return keys, index


def solve2(data):
    triple_regex = re.compile(r'(.)\1\1')
    calc = Calculator(data)
    keys = []
    index = -1
    while len(keys) < 64:
        index += 1
        md5 = calc.get_stretched_hash(index)
        match = triple_regex.search(md5)
        if match:
            subindex = index + 1
            seq = match.group(1) * 5
            if any(seq in calc.get_stretched_hash(subindex + i) for i in range(1000)):
                print("found key {0} at index {1}".format(md5, index))
                keys.append(md5)
    return keys, index


if __name__ == '__main__':

    data = 'zpqevtbw'
    # data = 'abc'
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
