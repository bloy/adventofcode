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

    def shift_letter(self, letter):
        if letter == '-':
            return ' '
        return chr((ord(letter) - ord('a') + int(self.sector)) % 26 + ord('a'))

    @property
    def decrypted_name(self):
        return "".join(self.shift_letter(letter) for letter in self.name)


def solve1(data):
    rooms = (Room.unpack(item) for item in data)
    return sum(int(room.sector) for room in rooms if room.valid)


def solve2(data):
    rooms = (Room.unpack(item) for item in data)
    return [room for room in rooms
            if room.valid and room.decrypted_name == 'northpole object storage']


if __name__ == '__main__':
    # data = [
        # 'aaaaa-bbb-z-y-x-123[abxyz]',
        # 'a-b-c-d-e-f-g-h-987[abcde]',
        # 'not-a-real-room-404[oarel]',
        # 'totally-a-real-room-200[decoy]',
    # ]
    with open('day4_input.txt') as f:
        data = [line for line in f.read().split("\n") if line]
    data = [d for d in data if d]

    pprint.pprint(solve1(data))
    pprint.pprint(solve2(data))
