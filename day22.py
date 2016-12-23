#!/usr/bin/env python3

import collections
import itertools
import pprint
import re

class Disc(collections.namedtuple('Disc', 'position size used avail')):
    __slots__ = ()

    @property
    def x(self):
        return self.position[0]

    @property
    def y(self):
        return self.position[1]

    @property
    def filesystem(self):
        return '/dev/grid/node-x{0}-y{1}'.format(*self.position)

    @property
    def map_mark(self):
        if self.used == 0:
            return '_ '
        if self.used > 100:
            return '# '
        return '. '

    def __repr__(self):
        return "Disc(position={0}, size={1}, used={2}, avail={3}, is_goal={4})".format(
            repr(self.position), repr(self.size), repr(self.used), repr(self.avail), repr(self.is_goal))

    def empty(self):
        return self.__class__(self.position, self.size, 0, self.size)

    def fill(self, used, is_goal):
        return self.__class__(self.position, self.size, used, self.size - used)

    @classmethod
    def from_line(cls, line):
        rgx = (r'.*node-x(?P<x>\d+)-y(?P<y>\d+)\s+(?P<size>\d+)T\s+'
               r'(?P<used>\d+)T\s+(?P<avail>\d+)T\s+(?P<percent>\d+)%\s*')
        match = re.match(rgx, line)
        groups = match.groupdict()
        return cls(position=(int(groups['x']), int(groups['y'])), size=int(groups['size']),
                   used=int(groups['used']), avail=int(groups['avail']))


def viable_pairs(discs):
    for a, b in itertools.combinations(discs, 2):
        if ((a.used != 0 and a.used <= b.avail) or (b.used != 0 and b.used <= a.avail)):
            yield (a, b)


def num_viable_pairs(discs):
    count = 0
    for pair in viable_pairs(discs):
        count += 1
    return count



class DiscArray(object):
    def __init__(self, data):
        self.discs = [
            list(group) for i, group in itertools.groupby(
                sorted(data, key=lambda x: (x.position[1], x.position[0])),
                key=lambda x: x.position[1])]
        self.max_x = len(self.discs[0])-1
        self.max_y = len(self.discs)-1
        self.access_point = (0, 0)
        self.goal = (self.max_x, 0)

    def __str__(self):
        return "\n".join(
            "".join(self.map_mark(disc) for disc in row)
            for row in self.discs)

    def map_mark(self, disc):
        if disc.position == self.access_point:
            return '! '
        elif disc.position == self.goal:
            return 'G '
        else:
            return disc.map_mark


if __name__ == '__main__':
    with open('day22_input.txt') as f:
        data = tuple([Disc.from_line(line) for line in f.readlines()
                if line and 'node' in line])

    print(num_viable_pairs(data))
    discs = DiscArray(data)
    print(discs)
