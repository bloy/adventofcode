#!env python

import aoc
import collections
import itertools
import pprint


def flip(chunk, vertical=True):
    size = int(len(chunk) ** 0.5)
    if vertical:
        return "".join(reversed(list(chunk[x*size:x*size+size]
                                for x in range(size))))
    return "".join("".join(reversed(chunk[x*size:x*size+size]))
                   for x in range(size))

def rotate(chunk):
    size = int(len(chunk) ** 0.5)
    return "".join("".join(reversed(chunk[x:9:size])) for x in range(size))


def rotates_and_flips(chunk):
    trans = set()
    trans.add(rotate(chunk))
    trans.add(rotate(rotate(chunk)))
    trans.add(rotate(rotate(rotate(chunk))))
    trans.add(flip(chunk))
    trans.add(rotate(flip(chunk)))
    trans.add(rotate(rotate(flip(chunk))))
    trans.add(rotate(rotate(rotate(flip(chunk)))))
    trans.add(flip(chunk, False))
    trans.add(rotate(flip(chunk, False)))
    trans.add(rotate(rotate(flip(chunk, False))))
    trans.add(rotate(rotate(rotate(flip(chunk, False)))))
    return tuple(trans)


def split_parts(large):
    bigsize = int(len(large) ** 0.5)
    if bigsize % 2 == 0:
        chunksize = 2
    else:
        chunksize = 3
    parts = []
    partsize = int(bigsize / chunksize)
    for y in range(partsize):
        for x in range(partsize):
            startx = x * chunksize
            starty = y * chunksize
            part = "".join(
                large[startx+bigsize*i:startx+bigsize*i+chunksize]
                for i in range(starty, starty+chunksize))
            parts.append(part)
    return parts


def join_parts(parts):
    partsize = int(len(parts) ** 0.5)
    chunksize = int(len(parts[0]) ** 0.5)
    parts = [[part[i*chunksize:i*chunksize+chunksize]
              for i in range(chunksize)]
             for part in parts]
    rows = []
    for i in range(partsize):
        subrows = parts[i*partsize:i*partsize+partsize]
        subrows = ["".join(subrow) for subrow in zip(*subrows)]
        rows.append(subrows)
    return "".join("".join(row) for row in rows)


def parse_data(lines):
    transforms = {}
    for line in lines:
        ingrid, outgrid = ["".join(part.split('/'))
                           for part in line.split(' => ')]
        transforms[ingrid] = outgrid
    for key, value in tuple(transforms.items()):
        for change in rotates_and_flips(key):
            if change not in transforms:
                transforms[change] = value
    return transforms


def solve1(start, data):
    canvas = start
    print_chunk(canvas)
    for i in range(5):
        parts = split_parts(canvas)
        parts = [data[part] for part in parts]
        canvas = join_parts(parts)
        print_chunk(canvas)
    return canvas.count('#')


def solve2(start, data):
    canvas = start
    print_chunk(canvas)
    for i in range(18):
        parts = split_parts(canvas)
        parts = [data[part] for part in parts]
        canvas = join_parts(parts)
        print_chunk(canvas)
    return canvas.count('#')


def print_chunk(chunk):
    size = int(len(chunk) ** 0.5)
    print("\n".join(chunk[x*size:x*size+size] for x in range(size)))


lines = [
    '../.# => ##./#../...',
    '.#./..#/### => #..#/..../..../#..#',
]

start = (
    '.#.'
    '..#'
    '###'
)

test = (
    '01234567'
    '89abcdef'
    'ghijklmn'
    'opqrstuv'
    'wxyzABCD'
    'EFGHIJKL'
    'MNOPQRST'
    'UVWXYZ*$'
)

if __name__ == '__main__':
    lines = aoc.input_lines(day=21)
    data = parse_data(lines)
    pprint.pprint(solve1(start, data))
    pprint.pprint(solve2(start, data))
