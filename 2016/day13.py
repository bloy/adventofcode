#!/usr/bin/env python3

import collections


def is_space(input_number, x, y):
    if x < 0 or y < 0:
        return False
    num = x*x + 3*x + 2*x*y + y + y*y + input_number
    num_ones = len([c for c in "{0:b}".format(num) if c == "1"])
    return ((num_ones % 2) == 0)

def print_map(input_number, size):
    map_chars = {
        True: '.',
        False: '#'
    }
    for y in range(size):
        for x in range(size):
            if is_space(input_number, x, y):
                c = '.'
            else:
                c = '#'
            print(c, end='')
        print()


def find_path(input_number, start, goal):
    seen = set()

    queue = collections.deque()
    queue.append((start, tuple()))

    while queue:
        step = queue.popleft()
        position = step[0]
        path = step[1]
        if position == goal:
            return path
        if position not in seen:
            seen.add(position)
            x = position[0]
            y = position[1]
            for p in ((x-1, y), (x, y-1), (x+1, y), (x, y+1)):
                if p not in seen and is_space(input_number, p[0], p[1]):
                    queue.append((p, path + (position,)))


def find_locations(input_number, start, num_steps):
    seen = set()

    queue = collections.deque()
    queue.append((start, 0))

    while queue:
        step = queue.popleft()
        position = step[0]
        length = step[1]
        if position not in seen:
            seen.add(position)
            if length < num_steps:
                x = position[0]
                y = position[1]
                for p in ((x-1, y), (x, y-1), (x+1, y), (x, y+1)):
                    if p not in seen and is_space(input_number, p[0], p[1]):
                        queue.append((p, length + 1))
    return seen


if __name__ == '__main__':
    start = (1, 1)
    input_number = 1358
    goal = (31, 39)

    # input_number = 10
    # goal = (7, 4)

    print_map(input_number, max(goal) + 1)
    path = find_path(input_number, start, goal)
    print(path)
    print(len(path))

    seen = find_locations(input_number, start, 50)
    print()
    print(seen)
    print(len(seen))

