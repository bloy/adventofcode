#!/bin/env python3

import collections
import hashlib
import pprint


def md5(indata):
    hasher = hashlib.md5()
    hasher.update(indata.encode('utf-8'))
    return hasher.hexdigest()


class State(collections.namedtuple('State', ('position', 'path'))):
    __slots__ = ()

    def valid_move(self, direction):
        return ((direction == 'U' and self.position[1] > 0) or
                (direction == 'D' and self.position[1] < 3) or
                (direction == 'L' and self.position[0] > 0) or
                (direction == 'R' and self.position[0] < 3))

    def move_position(self, direction):
        if direction == 'U':
            return (self.position[0], self.position[1] - 1)
        elif direction == 'D':
            return (self.position[0], self.position[1] + 1)
        elif direction == 'L':
            return (self.position[0] - 1, self.position[1])
        elif direction == 'R':
            return (self.position[0] + 1, self.position[1])

    def next_moves(self, passcode):
        path_hash = md5(passcode + self.path)
        return tuple([
            State(position=self.move_position(direction),
                  path=self.path+direction)
            for i, direction in enumerate('UDLR')
            if self.valid_move(direction) and path_hash[i] in 'bcdef'])


def solve1(passcode):
    initial_state = State(position=(0, 0), path='')
    goal = (3, 3)
    pending = collections.deque()
    pending.append(initial_state)
    node_count = 0

    while pending:
        node_count += 1
        state = pending.popleft()
        position = state.position
        if position == goal:
            return node_count, state.path
        for move in state.next_moves(passcode):
            pending.append(move)

def solve2(passcode):
    initial_state = State(position=(0, 0), path='')
    goal = (3, 3)
    pending = collections.deque()
    pending.append(initial_state)

    longest_path = ''
    node_count = 0

    print("")

    while pending:
        node_count += 1
        state = pending.pop()
        position = state.position
        if position == goal:
            if len(longest_path) < len(state.path):
                longest_path = state.path
        else:
            for move in state.next_moves(passcode):
                pending.append(move)

    return node_count, len(longest_path)



if __name__ == '__main__':
    passcode = 'ioramepc'
    # passcode = 'ulqzkmiv'
    pprint.pprint(solve1(passcode))
    pprint.pprint(solve2(passcode))
