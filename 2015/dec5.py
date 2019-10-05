#!/usr/bin/env python
import re

BAD_SUBSTRINGS = ['ab', 'cd', 'pq', 'xy']
VOWELS = set('aeiou')

def is_nice1(string):
    if any([string.find(s) >= 0 for s in BAD_SUBSTRINGS]):
        return False
    if not re.search(r'(.)\1', string):
        return False
    if len([c for c in string if c in VOWELS]) < 3:
        return False
    return True


def is_nice2(string):
    if re.search(r'(.)(.).*\1\2', string) and re.search(r'(.).\1', string):
        return True
    else:
        return False


if __name__ == '__main__':
    test_strings1 = (
        ('ugknbfddgicrmopn', True),
        ('aaa', True),
        ('jchzalrnumimnmhp', False),
        ('haegwjzuvuyypxyu', False),
        ('dvszwmarrgswjxmb', False)
    )
    test_strings2 = (
        ('qjhvhtzxzqqjkmpb', True),
        ('xxyxx', True),
        ('uurcxstgmygtbstg', False),
        ('ieodomkazucvgmuy', False)
    )
    for (string, expected) in test_strings1:
        print(string, expected, is_nice1(string))
    for (string, expected) in test_strings2:
        print(string, expected, is_nice2(string))

    with open('input/day_5') as lines:
        nice_count1 = 0
        naughty_count1 = 0
        nice_count2 = 0
        naughty_count2 = 0
        for line in lines:
            if is_nice1(line):
                nice_count1 += 1
            else:
                naughty_count1 += 1
            if is_nice2(line):
                nice_count2 += 1
            else:
                naughty_count2 += 1
        print("Nice count 1:", nice_count1)
        print("Naughty count 1:", naughty_count1)
        print("Nice count 2:", nice_count2)
        print("Naughty count 2:", naughty_count2)
