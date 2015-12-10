#!/usr/bin/env python
import itertools

def look_and_say(input_str):
    output = []
    for char, char_list in itertools.groupby(input_str):
        count = len(list(char_list))
        output += str(count)
        output += char
    return "".join(output)

def run_part1(start_string):
    current_string = start_string
    for x in range(41):
        length = len(current_string)
        if length > 80:
            print (x, ": big :", length)
        else:
            print(x, ":", current_string, ":", length)
        current_string = look_and_say(current_string)


if __name__ == '__main__':
    DAY_10_INPUT = '3113322113'
    run_part1(DAY_10_INPUT)
