#!env python

import aoc


def iterate(row):
    banks = list(row)
    idx = banks.index(max(banks))
    count = banks[idx]
    banks[idx] = 0
    while count > 0:
        idx = (idx + 1) % len(banks)
        banks[idx] += 1
        count -= 1
    return tuple(banks)


def solve1(row):
    seen = set()
    count = 0
    while row not in seen:
        seen.add(row)
        count += 1
        row = iterate(row)
    return count

def solve2(row):
    seen = set()
    record = list()
    count = 0
    while row not in seen:
        seen.add(row)
        record.append(row)
        count += 1
        row = iterate(row)
    return count - record.index(row)

if __name__ == '__main__':
    rows = [(0, 2, 7, 0)]
    rows = [tuple(int(bank) for bank in row.split()) for row in aoc.input_lines(day=6)]
    row = rows[0]
    print(solve1(row))
    print(solve2(row))
