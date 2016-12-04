#!/bin/env python3

from __future__ import unicode_literals
import collections
import itertools
import pprint
import re

class Room(collections.namedtuple('Room', 'name sector checksum')):
    __slots__ = ()
    roomexp = re.compile(
        r'^(?P<name>[\w-]+)-(?P<sector>\d+)\[(?P<checksum>\w+)\]$')

    @classmethod
    def unpack(cls, name):
        match = cls.roomexp.match(name)
        return cls(**match.groupdict()) if match else None

    @property
    def computed_checksum(self):
        counts = collections.Counter(self.name.replace('-', ''))
        count_groups = itertools.groupby(
            counts.most_common(), lambda count: count[1])
        letters = (item[0]
                   for group in count_groups
                   for item in sorted(group[1], key=lambda x: x[0]))
        return "".join(itertools.islice(letters, 5))

    @property
    def valid(self):
        return self.computed_checksum == self.checksum


def solve1(data):
    rooms = (Room.unpack(item) for item in data)
    return sum(int(room.sector) for room in rooms if room.valid)



def solve2(data):
    rooms = (Room.unpack(item) for item in data)


if __name__ == '__main__':
    data = [
        'aaaaa-bbb-z-y-x-123[abxyz]',
        'a-b-c-d-e-f-g-h-987[abcde]',
        'not-a-real-room-404[oarel]',
        'totally-a-real-room-200[decoy]',
    ]

    print(solve1(data))
    print(solve2(data))
