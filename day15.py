#!/bin/env python3

import collections
import itertools
import pprint


class Disc(collections.namedtuple('Disc', ('num', 'positions', 'start'))):
    @property
    def positions_by_time(self):
        return (((x + self.num + self.start) % self.positions)
                for x in itertools.count(start=0))


def solve1(discs):
    position_generators = [disc.positions_by_time for disc in discs]
    for x, positions in enumerate(zip(*position_generators)):
        if all([position == 0 for position in positions]):
            return x


def solve2(discs):
    position_generators = [disc.positions_by_time for disc in discs]
    for x, positions in enumerate(zip(*position_generators)):
        if all([position == 0 for position in positions]):
            return x


if __name__ == '__main__':

    discs = (
        Disc(1, 5, 2),
        Disc(2, 13, 7),
        Disc(3, 17, 10),
        Disc(4, 3, 2),
        Disc(5, 19, 9),
        Disc(6, 7, 0),
    )

    discs_part2 = discs + (Disc(7, 11, 0), )

    # discs = (
        # Disc(1, 5, 4),
        # Disc(2, 2, 1),
    # )
    pprint.pprint(solve1(discs))
    pprint.pprint(solve2(discs_part2))
