#!/usr/bin/env python

with open('input/day_1') as infile:
    data = infile.read()

floor = 0
for c in data:
    if c == '(':
        floor += 1
    elif c == ')':
        floor -= 1

print(floor)
