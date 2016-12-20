#!/usr/bin/env python3

def solve1(data):
    test_num = 0
    prv_rng = range(0, 0)
    for rng in sorted(data, key=lambda x: x.start):
        print("testing {0} against range {1}, {2} (previous {3}, {4})".format(
            test_num, rng.start, rng.stop, prv_rng.start, prv_rng.stop))
        if rng.stop in prv_rng:
            continue
        if test_num in rng or test_num == rng.stop:
            prv_rng = rng
            test_num = rng.stop + 1
        else:
            return test_num

def solve2(data):
    return 0

def parse_line(line):
    (start, end) = line.split('-')
    return range(int(start), int(end))

if __name__ == '__main__':
    with open('day20_input.txt') as f:
        data = [parse_line(line) for line in f.readlines() if line]

    print(solve1(data))
    print(solve2(data))
