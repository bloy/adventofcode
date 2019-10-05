#!/usr/bin/env python
import math

def factors(number):
    divisors = [x for x in range(1, int(math.sqrt(number)) + 1)
                if number % x == 0]
    large = [number // divisor for divisor in divisors
             if number != divisor * divisor]
    return divisors + large


def present_count(house_number):
    elves = factors(house_number)
    part1_presents = sum(elves) * 10
    part2_presents = sum(elf for elf in elves if elf * 50 >= house_number) * 11
    return (part1_presents, part2_presents)


if __name__ == '__main__':
    min_presents = 36000000
    part1 = None
    part2 = None
    house_number = 0
    while part1 is None or part2 is None:
        house_number += 1
        presents = present_count(house_number)
        if house_number % 10000 == 0:
            print(house_number, ":", presents)
        if presents[0] >= min_presents and part1 is None:
            part1 = (house_number, presents[0])
            print("part 1", part1)
        if presents[1] >= min_presents and part2 is None:
            part2 = (house_number, presents[1])
            print("part 2", part2)
    print("part 1", part1)
    print("part 2", part2)
