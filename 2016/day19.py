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


def solve2(count):
    """Pattern is similar for part 2.

    count = n
    n = 3**a + l + m  (where l is 3**a or 0)
    f(n) = m iff l is 0
    f(n) = 2m + l iff l is 3**a
    """
    power = 0
    while 3 ** power < count:
        power += 1
    a = power - 1
    m = count - 3**a
    if m > 3**a:
        return 2 * (m - 3**a) + 3**a
    else:
        return m

if __name__ == '__main__':
    count = 3014603
    # count = 5
    print(solve1(count))
    print(solve2(count))
