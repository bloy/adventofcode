#!/usr/bin/env python3

import itertools
import hashlib
import pprint


def md5(indata):
    hasher = hashlib.md5()
    hasher.update(indata.encode('utf-8'))
    return hasher.hexdigest()


def solve1(data):
    generator = filter(lambda x: x[0:5] == '00000',
                       (md5(data + str(num)) for num in itertools.count()))
    return "".join(hashstr[5] for hashstr in itertools.islice(generator, 8))


def solve2(data):
    return 0


if __name__ == '__main__':
    sample_data = 'abc'
    data = 'ugkcyxxp'

    pprint.pprint(solve1(sample_data))
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(sample_data))
    pprint.pprint(solve2(data))
