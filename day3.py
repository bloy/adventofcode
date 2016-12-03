#!/bin/env python

from __future__ import unicode_literals


def valid_triangle(tri):
    return (
        tri[0] + tri[1] > tri[2] and
        tri[1] + tri[2] > tri[0] and
        tri[2] + tri[0] > tri[1])


def solve1(lines):
    triangles = [tuple(int(num) for num in line.split())
                 for line in lines.split("\n") if line != '']
    return len([tri for tri in triangles if valid_triangle(tri)])



if __name__ == '__main__':
    with open('day3_input.txt') as f:
        lines = f.read()

    print("number of valid horizontal triangles: {0}".format(solve1(lines)))
