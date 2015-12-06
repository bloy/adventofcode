#!/usr/bin/env python
import re


class LightGrid(object):
    X_SIZE = 1000
    Y_SIZE = 1000
    INSTRUCTION_REGEX = re.compile(r'^(toggle|turn (on|off)) '
                                   r'(\d+),(\d+) through (\d+),(\d+)')
    def __init__(self):
        print("initializing light grid")
        self._lights = {}
        for x in range(self.X_SIZE):
            for y in range(self.Y_SIZE):
                self._lights[(x,y)] = False

    def toggle(self, position):
        self._lights[position] = not self._lights[position]

    def turn_on(self, position):
        self._lights[position] = True

    def turn_off(self, position):
        self._lights[position] = False

    def run_instruction(self, line):
        print("running instruction: {0}".format(line.strip()))
        groups = self.INSTRUCTION_REGEX.match(line.strip()).groups()
        instruction = groups[0]
        start_pos = (int(groups[2]), int(groups[3]))
        end_pos = (int(groups[4]), int(groups[5]))
        if instruction == 'toggle':
            function = self.toggle
        elif instruction == 'turn on':
            function = self.turn_on
        elif instruction == 'turn off':
            function = self.turn_off
        for x in range(start_pos[0], end_pos[0]+1):
            for y in range(start_pos[1], end_pos[1]+1):
                function((x,y))

    def lit_lights(self):
        return [self._lights[key]
                for key in self._lights.keys()
                if self._lights[key]]

    def lit_count(self):
        return len(self.lit_lights())


if __name__ == '__main__':
    grid = LightGrid()
    with open('input/day_6') as input_file:
        for line in input_file:
            grid.run_instruction(line)

    print("Lit light count:", grid.lit_count())
