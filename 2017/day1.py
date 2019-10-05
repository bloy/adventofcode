#!env python
import aoc

def revcaptcha(rotate, seed):
    return sum(int(digit) for i, digit in enumerate(seed)
               if seed[(i+rotate) % len(seed)] == digit)

if __name__ == '__main__':
    puzzle_input = "".join(aoc.input_lines(day=1))

    rotate1 = 1
    rotate2 = len(puzzle_input) // 2

    print(revcaptcha(rotate1, puzzle_input))
    print(revcaptcha(rotate2, puzzle_input))
