#/usr/bin/env python
import collections

with open('input/day_3') as file:
    data = file.read()

houses = collections.Counter()
position_x = [0, 0]
position_y = [0, 0]
houses[(0,0)] = 2
santa = 0
for direction in data.strip():
    if direction == '^':
        position_y[santa] += 1
    elif direction == 'v':
        position_y[santa] -= 1
    elif direction == '>':
        position_x[santa] += 1
    elif direction == '<':
        position_x[santa] -= 1
    else:
        print('unknown character "{0}"'.format(direction))
    houses[(position_x[santa], position_y[santa])] += 1
    if santa == 1:
        santa = 0
    else:
        santa = 1

print(len(houses.keys()))
