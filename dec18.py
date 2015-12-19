#!/usr/bin/env python
import pprint
from collections import Counter

NEIGHBORS = {(x, y) for y in range(-1, 2) for x in range(-1, 2) if x or y}

def parse_lines(lines):
    board = {
        (x, y) for y, line in enumerate(lines)
        for x, char in enumerate(line)
        if char == '#'
    }
    return board

def neighbors(point, size):
    return {(x + point[0], y + point[1]) for (x, y) in NEIGHBORS
     if (x + point[0] >= 0 and x + point[0] < size and
         y + point[1] >= 0 and y + point[1] < size)}

def step(board, size):
    next_step = Counter(neighbor for cell in board
                     for neighbor in neighbors(cell, size))
    return {cell for cell in next_step
            if next_step[cell] == 3 or
            (next_step[cell] == 2 and cell in board)}

if __name__ == '__main__':
    with open('input/day_18') as lines:
        board = parse_lines(lines)
    size = 100
    for x in range(100):
        board = step(board, size)
    pprint.pprint(board)
    pprint.pprint(len(board))
