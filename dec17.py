#!/usr/bin/env python
import itertools
import collections


def find_possible_containers(target, containers):
    counters = collections.Counter()
    for number in range(1, len(containers)+1):
        for possible in itertools.combinations(containers, number):
            if sum(possible) == target:
                counters[number] += 1
    return counters


if __name__ == '__main__':
    containers = [43, 3, 4, 10, 21, 44, 4, 6, 47, 41,
                  34, 17, 17, 44, 36, 31, 46, 9, 27, 38]
    volume = 150

    possible_counts = find_possible_containers(volume, containers)
    for number in sorted(possible_counts.keys()):
        print(number, possible_counts[number])
    print(sum(possible_counts.values()))
