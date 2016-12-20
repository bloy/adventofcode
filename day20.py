#!/usr/bin/env python3

def solve(data):
    allowed = []
    test_num = 0
    prv_rng = range(0, 0)
    for rng in sorted(data, key=lambda x: x.start):
        if rng.stop in prv_rng:
            continue
        while not test_num in rng:
            allowed.append(test_num)
            test_num += 1
        prv_rng = rng
        test_num = rng.stop
    return allowed


def parse_line(line):
    (start, end) = line.split('-')
    return range(int(start), int(end) + 1)

if __name__ == '__main__':
    with open('day20_input.txt') as f:
        data = [parse_line(line) for line in f.readlines() if line]

    allowed = solve(data)
    print(allowed[0])
    print(len(allowed))

