#!/usr/bin/env python

with open('input/day_2') as f:
    paper_total = 0
    ribbon_total = 0
    for line in f:
        edges = [int(n) for n in line.strip().split('x')]
        sides = (edges[0] * edges[1],
                 edges[1] * edges[2],
                 edges[0] * edges[2])
        perimeter = 2 * min(edges[0] + edges[1],
                            edges[1] + edges[2],
                            edges[0] + edges[2])
        volume = edges[0] * edges[1] * edges[2]
        minside = min(sides)
        paper_total += (sides[0] + sides[1] + sides[2]) * 2 + minside
        ribbon_total += perimeter + volume
    print(paper_total)
    print(ribbon_total)
