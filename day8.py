#!/usr/bin/env python3

import pprint
import re


COMMANDS = {
    'rect': re.compile(r'^rect (\d+)x(\d+)'),
    'rcol': re.compile(r'^rotate column x=(\d+) by (\d+)'),
    'rrow': re.compile(r'^rotate row y=(\d+) by (\d+)'),
}


class Screen(object):

    def __init__(self, x_size, y_size):
        self.pixels = [False for y in range(y_size) for x in range(x_size)]
        self.x_size = x_size
        self.y_size = y_size

    def __str__(self):
        return "\n".join(
            "".join(
                "#" if self.pixels[y*self.x_size + x] else "."
                for x in range(self.x_size))
            for y in range(self.y_size))

    def set_pixel(self, x, y, value):
        self.pixels[y*self.x_size + x] = value


    def rect(self, x, y):
        for ypos in range(y):
            for xpos in range(x):
                self.set_pixel(xpos, ypos, True)

    def rotatecol(self, col, by):
        self.pixels[col::self.x_size] = (
            self.pixels[col::self.x_size][-by:] +
            self.pixels[col::self.x_size][:-by])

    def rotaterow(self, row, by):
        self.pixels[row*self.x_size:((row+1)*self.x_size)] = (
            self.pixels[row*self.x_size:((row+1)*self.x_size)][-by:] +
            self.pixels[row*self.x_size:((row+1)*self.x_size)][:-by])

    def num_lit(self):
        return len(list(pixel for pixel in self.pixels if pixel))


def convert_line(line):
    for command in COMMANDS:
        match = COMMANDS[command].match(line)
        if match:
            return (command, int(match.group(1)), int(match.group(2)))
    return None


def solve1(screen, data):
    for line in data:
        cmd = convert_line(line)
        if cmd[0] == 'rect':
            screen.rect(*cmd[1:])
        elif cmd[0] == 'rcol':
            screen.rotatecol(*cmd[1:])
        elif cmd[0] == 'rrow':
            screen.rotaterow(*cmd[1:])
        else:
            raise NotImplementedError("{0} not implemented".format(cmd[0]))
    print(screen)
    return screen.num_lit()


def solve2(screen, data):
    return solve1(screen, data)


if __name__ == '__main__':
    with open('day8_input.txt') as f:
        data = [line for line in f.read().split("\n") if line != '']

    x_size = 50
    y_size = 6

    # data = [
        # 'rect 3x2',
        # 'rotate column x=1 by 1',
        # 'rotate row y=0 by 4',
        # 'rotate column x=1 by 1',
        # 'rect 3x2'
    # ]

    # x_size = 7
    # y_size = 3

    print(solve1(Screen(x_size, y_size), data))
    print(solve2(Screen(x_size, y_size), data))
