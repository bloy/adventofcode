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
    generator = filter(lambda x: x[0:5] == '00000',
                       (md5(data + str(num)) for num in itertools.count()))
    value = list('________')
    for md5sum in generator:
        (position, replacement) = (md5sum[5], md5sum[6])
        if position in '01234567' and value[int(position)] == '_':
            value[int(position)] = replacement
        print("".join(value), position, replacement, end="\r")
        if '_' not in value:
            print()
            return "".join(value)


if __name__ == '__main__':
    sample_data = 'abc'
    data = 'ugkcyxxp'

    pprint.pprint(solve1(sample_data))
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(sample_data))
    pprint.pprint(solve2(data))
