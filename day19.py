#!/bin/env python3

def solve1(count):
    # https://en.wikipedia.org/wiki/Josephus_problem
    # https://www.youtube.com/watch?v=uCsD3ZGzMgE

    power = 0
    while 2 ** power < count:
        power += 1
    a = power - 1
    l = count - 2**a
    return a, l, 2 * l + 1


if __name__ == '__main__':
    count = 3014603
    # count = 5
    print(solve1(count))
