#/usr/bin/env python
import collections

with open('input/day_3') as file:
    data = file.read()

houses = collections.Counter()
position_x = 0
position_y = 0
houses[(0,0)] = 1
for direction in data.strip():
    if direction == '^':
        position_y += 1
    elif direction == 'v':
        position_y -= 1
    elif direction == '>':
        position_x += 1
    elif direction == '<':
        position_x -= 1
    else:
        print('unknown character "{0}"'.format(direction))
    houses[(position_x, position_y)] += 1

print(len(houses.keys()))
