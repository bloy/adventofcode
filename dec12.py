#!/usr/bin/env python
import re
import json


def find_numbers(in_data):
    numbers = []
    for line in in_data:
        numbers += [int(num) for num in re.findall(r'(-?\d+)', line)]
    print(numbers)
    return numbers

def no_red_accounting(data):
    if isinstance(data, int):
        return data
    if isinstance(data, list):
        return sum(no_red_accounting(value) for value in data)
    if isinstance(data, dict):
        if any(data[key] == 'red' for key in data.keys()):
            return 0
        return sum(no_red_accounting(data[key]) for key in data)
    return 0


if __name__ == '__main__':
    with open('input/day_12') as infile:
        print(sum(find_numbers(infile)))

    with open('input/day_12') as infile:
        data = json.load(infile)
        print(no_red_accounting(data))
