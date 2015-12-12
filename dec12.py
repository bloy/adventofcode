#!/usr/bin/env python
import re


def find_numbers(in_data):
    numbers = []
    for line in in_data:
        numbers += [int(num) for num in re.findall(r'(-?\d+)', line)]
    print(numbers)
    return numbers


if __name__ == '__main__':
    with open('input/day_12') as infile:
        print(sum(find_numbers(infile)))

