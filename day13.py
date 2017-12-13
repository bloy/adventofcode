#!env python

import aoc
import collections
import pprint


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


def position_at(traverse, time):
    offset = time % ((traverse - 1) * 2)
    return 2 * (traverse - 1) - offset if offset > traverse - 1 else offset


def simulate(layers, start=0):
    collisions = []
    for t in range(max(layers.keys()) + 1):
        if t in layers and position_at(layers[t], t+start) == 0:
            collisions.append((t, layers[t]))
    return sum((a * b) for a,b in collisions)


def print_at(layers, time):
    print("----------")
    print("Time:", time)
    for t in range(max(layers.keys()) + 1):
        if t in layers:
            print("{0}: {1}/{2}".format(t,
                                        position_at(layers[t], time),
                                        layers[t]))
        else:
            print("{0}: ---".format(t))

def solve1(data):
    layers = data
    return simulate(layers)


def solve2(data):
    t = 0
    layers = data
    caught = simulate(layers, start=t)
    while caught > 0:
        print(t, caught)
        t += 1
        caught = simulate(layers, start=t)
    return t



lines = [
    "0: 3",
    "1: 2",
    "4: 6",
    "6: 4",
]


if __name__ == '__main__':
    #lines = aoc.input_lines(day=13)
    data = parse_data(lines)
    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
    # for i in range(11):
        # print_at(data, i)
