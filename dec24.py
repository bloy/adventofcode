#!/usr/bin/env python
import itertools
import functools

PACKAGES = (
    1, 2, 3, 5, 7, 13, 17, 19, 23, 29, 31, 37, 41, 43, 53, 59, 61,
    67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
)

if __name__ == '__main__':
    total_sum = sum(PACKAGES)
    third_sum = total_sum / 3
    quarter_sum = total_sum / 4
    for num_packages in range(1, len(PACKAGES)):
        possibles = [group for group
                     in itertools.combinations(PACKAGES, num_packages)
                     if sum(group) == quarter_sum]
        if len(possibles) > 0:
            break
    bestqe = None
    for possible in possibles:
        qe = functools.reduce(lambda x,y: x*y, possible)
        if bestqe is None or qe < bestqe:
            bestqe = qe
    print(bestqe)
