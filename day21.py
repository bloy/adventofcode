#!/usr/bin/env python3

import itertools

def swap_position(word, x, y):
    word[x], word[y] = word[y], word[x]
    return word

def swap_letter(word, x, y):
    return swap_position(word, word.index(x), word.index(y))

def rotate_steps(word, direction, x):
    if direction == 'right':
        return word[-x:] + word[:-x]
    else:
        return word[x:] + word[:x]

def rotate_pos(word, x):
    index = word.index(x)
    word = rotate_steps(word, 'right', 1)
    word = rotate_steps(word, 'right', index)
    if index >= 4:
        word = rotate_steps(word, 'right', 1)
    return word

def reverse_pos(word, x, y):
    word[x:y+1] = reversed(word[x:y+1])
    return word

def move_pos(word, x, y):
    word[x:x+1], word[y:y] = [], word[x:x+1]
    return word

def parse_line(line):
    words = line.split()
    if words[0] == 'swap':
        if words[1] == 'position':
            return (swap_position, int(words[2]), int(words[5]))
        elif words[1] == 'letter':
            return (swap_letter, words[2], words[5])
    elif words[0] == 'rotate':
        if words[1] in ('left', 'right'):
            return (rotate_steps, words[1], int(words[2]))
        elif words[1] == 'based':
            return (rotate_pos, words[6])
    elif words[0] == 'reverse':
        return (reverse_pos, int(words[2]), int(words[4]))
    elif words[0] == 'move':
        return (move_pos, int(words[2]), int(words[5]))

def solve(word, data):
    word = list(word)
    for instruction in data:
        word = instruction[0](word, *instruction[1:])
    return "".join(word)

def reverse_solve(scrambled, data):
    for word in itertools.permutations(scrambled):
        if solve(word, data) == scrambled:
            return "".join(word)

if __name__ == '__main__':
    with open('day21_input.txt') as f:
        data = [parse_line(line) for line in f.readlines() if line]
    word = 'abcdefgh'

    word = solve(word, data)
    print(word)
    print(reverse_solve('fbgdceah', data))
