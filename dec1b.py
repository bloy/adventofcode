#!/usr/bin/env python

with open('input/day_1') as infile:
    data = infile.read()

floor = 0
first_basement = -1
for i, c in enumerate(data):
    if c == '(':
        floor += 1
    elif c == ')':
        floor -= 1
    if first_basement == -1 and floor < 0:
        first_basement = i+1

print(floor)
print(first_basement)
