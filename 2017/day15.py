#!env python

import pprint


CONSTA = 16807
CONSTB = 48271
DIVISOR = 2147483647


def parse_data(lines):
    scanners = {}
    maxlayer = 0
    for line in lines:
        layer, traverse = line.split(': ')
        layer = int(layer)
        traverse = int(traverse)
        if maxlayer < layer:
            maxlayer = layer
        scanners[layer] = traverse
    return scanners


def generators(starta, startb):
    numa = starta
    numb = startb
    bits16 = 2 ** 16 - 1
    while 1:
        numa = (numa * CONSTA) % DIVISOR
        numb = (numb * CONSTB) % DIVISOR
        yield (numa & bits16, numb & bits16)


def solve1(starta, startb, rounds):
    count = 0
    for i, values in enumerate(generators(starta, startb)):
        if values[1] == values[0]:
            count += 1
        if i < 10:
            print(i, values)
        if i % 1000 == 0:
            print(i, values)
        if rounds == i + 1:
            return count


def generator2(start, const, multiple):
    num = start
    bits16 = 2 ** 16 - 1
    while 1:
        num = (num * const) % DIVISOR
        if num % multiple == 0:
            yield num & bits16


def solve2(starta, startb, rounds):
    count = 0
    generatora = generator2(starta, CONSTA, 4)
    generatorb = generator2(startb, CONSTB, 8)
    for i, values in enumerate(zip(generatora, generatorb)):
        if i < 10 or i % 1000 == 0:
            print(i, count, values)
        if values[0] == values[1]:
            count += 1
        if rounds == i + 1:
            return count


starta = 512
startb = 191
rounds1 = 40 * 1000 * 1000
rounds2 = 5 * 1000 * 1000
# starta = 65   # example values
# startb = 8921
# rounds = 1057

if __name__ == '__main__':
    pprint.pprint(solve1(starta, startb, rounds1))
    pprint.pprint(solve2(starta, startb, rounds2))
