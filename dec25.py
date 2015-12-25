#!/usr/bin/env python

FIRST_VALUE = 20151125
INPUT_COLUMN = 3083
INPUT_ROW = 2978

def sequence_number(row, column):
    return sum(range(row + column - 1)) + column


if __name__ == '__main__':
    for x in range(1, 7):
        for y in range(1, 8 - x):
            print(x, y, sequence_number(x, y))
