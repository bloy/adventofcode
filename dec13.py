#!/usr/bin/env python
import re
import itertools
import pprint


def parse_lines(lines):
    parser_re = re.compile(r'^(?P<name1>\w+) would '
                           r'(?P<sign>gain|lose) (?P<number>\d+) happiness '
                           r'units by sitting next to (?P<name2>\w+)\.$')
    matrix = dict()
    guests = set()
    for line in lines:
        line = line.strip()
        groups = parser_re.match(line).groupdict()
        number = int(groups['number'])
        if groups['sign'] == 'lose':
            number = -1 * number
        matrix[(groups['name1'], groups['name2'])] = number
        guests.add(groups['name1'])
        guests.add(groups['name2'])
    return guests, matrix


def seating_combinations(seating):
    return itertools.chain(zip(seating, seating[1:] + (seating[0], )),
                           zip(seating[1:] + (seating[0], ), seating))


def seating_total(seating, matrix):
    return sum(matrix[pair] for pair in seating_combinations(seating))


def all_seatings(guests, matrix):
    return ((seating_total(seating, matrix), seating)
            for seating in itertools.permutations(guests))

def find_best(guests, matrix):
    best = None
    for candidate in all_seatings(guests, matrix):
        if best is None or best[0] < candidate[0]:
            best = candidate
    return best


def add_self_to_table(guests, matrix):
    for guest in guests:
        matrix[('me', guest)] = 0
        matrix[(guest, 'me')] = 0
    guests.add('me')
    return(guests, matrix)


if __name__ == '__main__':
    with open('input/day_13') as in_lines:
        guests, matrix = parse_lines(in_lines)
    best = find_best(guests, matrix)
    pprint.pprint(best)

    guests, matrix = add_self_to_table(guests, matrix)
    best_with_me = find_best(guests, matrix)
    pprint.pprint(best_with_me)
    print("difference:", best_with_me[0] - best[0])
