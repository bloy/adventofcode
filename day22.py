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

    def __repr__(self):
        return "Disc(position={0}, size={1}, used={2}, avail={3})".format(
            repr(self.position), repr(self.size), repr(self.used), repr(self.avail))

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
            print("{0}\n{1} is a viable pair\n".format(a,b))
            yield (a, b)


def num_viable_pairs(discs):
    count = 0
    for pair in viable_pairs(discs):
        count += 1
    return count


if __name__ == '__main__':
    with open('day22_input.txt') as f:
        data = tuple([Disc.from_line(line) for line in f.readlines()
                if line and 'node' in line])

    data = sorted(data, key=lambda x: x.position)
    pprint.pprint(data)
    #print(num_viable_pairs(data))
