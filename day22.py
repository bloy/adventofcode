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
        return cls(position=(groups['x'], groups['y']), size=groups['size'],
                   used=groups['used'], avail=groups['avail'])


if __name__ == '__main__':
    with open('day22_input.txt') as f:
        data = tuple([Disc.from_line(line) for line in f.readlines()
                if line and 'node' in line])

