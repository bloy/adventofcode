#!/usr/bin/env python

FIRST_VALUE = 20151125
INPUT_COLUMN = 3083
INPUT_ROW = 2978
MAGIC_MULTIPLIER = 252533
MAGIC_MOD = 33554393

def sequence_number(row, column):
    return sum(range(row + column - 1)) + column

def next_code(code):
    return (code * MAGIC_MULTIPLIER) % MAGIC_MOD

if __name__ == '__main__':
    code_sequence_number = sequence_number(INPUT_ROW, INPUT_COLUMN)
    current_code = FIRST_VALUE
    for i in range(code_sequence_number - 1):
        current_code = next_code(current_code)
    print(current_code)
