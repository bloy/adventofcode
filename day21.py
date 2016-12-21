#!/usr/bin/env python3

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
    print(word)
    word = list(word)
    for instruction in data:
        print(instruction[0].__name__, instruction[1:])
        word = instruction[0](word, *instruction[1:])
        print(word)
    return "".join(word)

if __name__ == '__main__':
    with open('day21_input.txt') as f:
        data = [parse_line(line) for line in f.readlines() if line]
    word = 'abcdefgh'

    # data = [
        # (swap_position, 4, 0),
        # (swap_letter, 'd', 'b'),
        # (reverse_pos, 0, 4),
        # (rotate_steps, 'left', 1),
        # (move_pos, 1, 4),
        # (move_pos, 3, 0),
        # (rotate_pos, 'b'),
        # (rotate_pos, 'd'),
    # ]
    # word = 'abcde'

    word = solve(word, data)
    print(word)
