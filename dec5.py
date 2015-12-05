#!/usr/bin/env python
import re

BAD_SUBSTRINGS = ['ab', 'cd', 'pq', 'xy']
VOWELS = set('aeiou')

def is_nice(string):
    if any([string.find(s) >= 0 for s in BAD_SUBSTRINGS]):
        return False
    if not re.search(r'(.)\1', string):
        return False
    if len([c for c in string if c in VOWELS]) < 3:
        return False
    return True


if __name__ == '__main__':
    test_strings = (
        ('ugknbfddgicrmopn', True),
        ('aaa', True),
        ('jchzalrnumimnmhp', False),
        ('haegwjzuvuyypxyu', False),
        ('dvszwmarrgswjxmb', False)
    )
    for (string, expected) in test_strings:
        print(string, expected, is_nice(string))

    with open('input/day_5') as lines:
        nice_count = 0
        naughty_count = 0
        for line in lines:
            if is_nice(line):
                nice_count += 1
            else:
                naughty_count += 1
        print("Nice count:", nice_count)
        print("Naughty count:", naughty_count)
