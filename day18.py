#!/bin/env python3

import pprint


def tiles(row):
    fake_row = '.' + row + '.'
    return "".join([('.' if fake_row[i] == fake_row[i+2] else '^')
                    for i in range(len(row))])


def solve(row, room_size):
    rows = [row]
    for i in range(room_size-1):
        row = tiles(row)
        rows.append(row)
    pprint.pprint(rows)
    return sum(len([x for x in y if x == '.']) for y in rows)


if __name__ == '__main__':
    row = '^.^^^..^^...^.^..^^^^^.....^...^^^..^^^^.^^.^^^^^^^^.^^.^^^^...^^...^^^^.^.^..^^..^..^.^^.^.^.......'
    room_size = 40
    # row = '.^^.^.^^^^'
    # room_size = 10
    print(solve(row, room_size))
