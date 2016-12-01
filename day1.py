#!/usr/bin/env python

from __future__ import unicode_literals

import collections

input_text = """R3, R1, R4, L4, R3, R1, R1, L3, L5, L5, L3, R1, R4, L2, L1, R3, L3, R2, R1, R1, L5, L2, L1, R2, L4, R1, L2, L4, R2, R2, L2, L4, L3, R1, R4, R3, L1, R1, L5, R4, L2, R185, L2, R4, R49, L3, L4, R5, R1, R1, L1, L1, R2, L1, L4, R4, R5, R4, L3, L5, R1, R71, L1, R1, R186, L5, L2, R5, R4, R1, L5, L2, R3, R2, R5, R5, R4, R1, R4, R2, L1, R4, L1, L4, L5, L4, R4, R5, R1, L2, L4, L1, L5, L3, L5, R2, L5, R4, L4, R3, R3, R1, R4, L1, L2, R2, L1, R4, R2, R2, R5, R2, R5, L1, R1, L4, R5, R4, R2, R4, L5, R3, R2, R5, R3, L3, L5, L4, L3, L2, L2, R3, R2, L1, L1, L5, R1, L3, R3, R4, R5, L3, L5, R1, L3, L5, L5, L2, R1, L3, L1, L3, R4, L1, R3, L2, L2, R3, R3, R4, R4, R1, L4, R1, L5"""

input_data = input_text.split(', ')

facing_machine = {
    'N': {'R': 'W', 'L': 'E'},
    'S': {'R': 'E', 'L': 'W'},
    'E': {'R': 'N', 'L': 'S'},
    'W': {'R': 'S', 'L': 'N'},
}

facings = {
    'N': complex(1, 0),
    'S': complex(-1, 0),
    'E': complex(0, 1),
    'W': complex(0, -1),
}

facing = 'N'
position = complex(0, 0)

visited = set()
visited.add(position)
first2 = None

for instruction in input_data:
    (turn, distance) = (instruction[0], int(instruction[1:]))
    facing = facing_machine[facing][turn]
    position = position + distance * facings[facing]
    if position in visited and first2 is None:
        first2 = position
    if position in visited:
        print("visited ({}, {}) again".format(position.real, position.imag))
    visited.add(position)

print(repr(position))
print("total taxicab distance: {}".format(abs(position.real) + abs(position.imag)))

print(repr(first2))
print("total taxicab distance: {}".format(abs(first2.real) + abs(first2.imag)))
