#!env python

import collections
import pprint


def solve1(states, steps):
    tape = collections.defaultdict(lambda: 0)
    cursor = 0
    state = 'A'
    percent = steps // 100
    for i in range(steps):
        curstate = states[state]
        curval = tape[cursor]
        tape[cursor] = curstate[curval]['write']
        cursor += curstate[curval]['move']
        state = curstate[curval]['next']
        if percent > 0 and i % percent == 0:
            print("step {0} of {1}: {2}%".format(
                i, steps, i * 100 / steps))
    return sum(tape.values())


LEFT = -1
RIGHT = 1
STATES = {
    'A': {
        0: {
            'write': 1,
            'move': RIGHT,
            'next': 'B',
        },
        1: {
            'write': 0,
            'move': RIGHT,
            'next': 'C',
        },
    },
    'B': {
        0: {
            'write': 0,
            'move': LEFT,
            'next': 'A',
        },
        1: {
            'write': 0,
            'move': RIGHT,
            'next': 'D',
        },
    },
    'C': {
        0: {
            'write': 1,
            'move': RIGHT,
            'next': 'D',
        },
        1: {
            'write': 1,
            'move': RIGHT,
            'next': 'A',
        },
    },
    'D': {
        0: {
            'write': 1,
            'move': LEFT,
            'next': 'E',
        },
        1: {
            'write': 0,
            'move': LEFT,
            'next': 'D',
        },
    },
    'E': {
        0: {
            'write': 1,
            'move': RIGHT,
            'next': 'F',
        },
        1: {
            'write': 1,
            'move': LEFT,
            'next': 'B',
        },
    },
    'F': {
        0: {
            'write': 1,
            'move': RIGHT,
            'next': 'A',
        },
        1: {
            'write': 1,
            'move': RIGHT,
            'next': 'E',
        },
    },
}

STEPS = 12368930

FAKESTATES = {
    'A': {
        0: {
            'write': 1,
            'move': RIGHT,
            'next': 'B',
        },
        1: {
            'write': 0,
            'move': LEFT,
            'next': 'B',
        },
    },
    'B': {
        0: {
            'write': 1,
            'move': LEFT,
            'next': 'A',
        },
        1: {
            'write': 1,
            'move': RIGHT,
            'next': 'A',
        },
    },
}
FAKESTEPS = 6


if __name__ == '__main__':
    pprint.pprint(solve1(FAKESTATES, FAKESTEPS))
    pprint.pprint(solve1(STATES, STEPS))
