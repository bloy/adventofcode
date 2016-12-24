#!/usr/bin/env python3

import collections
import itertools
import pprint

Space = collections.namedtuple('Space', 'x y char num')


class Maze(object):
    def __init__(self, lines):
        self.goals = []
        maze = []
        for y, line in enumerate(lines):
            row = []
            for x, char in enumerate(line):
                if char in '#.':
                    s = Space(x=x, y=y, char=char, num=None)
                else:
                    s = Space(x=x, y=y, char='.', num=char)
                    if char == '0':
                        self.start_pos = s
                    self.goals.append(s)
                row.append(s)
            maze.append(tuple(row))
        self.maze = tuple(maze)

    def maze_char(self, x, y):
        if self.maze[y][x].num:
            return self.maze[y][x].num
        else:
            return self.maze[y][x].char

    def __str__(self):
        return "\n".join("".join(self.maze_char(x, y)
                                 for x, c in enumerate(row))
                         for y, row in enumerate(self.maze))

    def next_moves(self, space):
        x = space.x
        y = space.y
        positions = ((x, y+1), (x, y-1), (x+1, y), (x-1, y))
        return [
            self.maze[p[1]][p[0]] for p in positions
            if self.maze[p[1]][p[0]].char == '.']

    def shortest_path(self, start, end):
        seen = set()
        pending = collections.deque()
        pending.append((start, tuple()))

        while pending:
            step = pending.popleft()
            position = step[0]
            path = step[1]
            if position == end:
                return len(path)
            if position not in seen:
                seen.add(position)
                for move in self.next_moves(position):
                    if move not in seen:
                        pending.append((move, path + (position, )))

    def solve_for_goals(self):
        self.distances = collections.defaultdict(dict)
        for start, end in itertools.combinations(self.goals, 2):
            length = self.shortest_path(start, end)
            self.distances[start.num][end.num] = length
            self.distances[end.num][start.num] = length
            print("the distance between {0} and {1} is {2}".format(
                start.num, end.num, length))

    def goal_path_length(self, path):
        distance = 0
        for pair in zip(path[0:-1], path[1:]):
            distance += self.distances[pair[0]][pair[1]]
        return distance


    def solve_for_distances(self, reset=False):
        start_goal = self.start_pos.num
        goal_set = set(goal.num for goal in self.goals if goal.num != start_goal)
        min_path = None
        min_length = None
        for path in itertools.permutations(goal_set):
            path = ('0',) + path
            if reset:
                path += ('0', )
            length = self.goal_path_length(path)
            print("length of path {0} is {1}".format(path, length))
            if not min_length or length < min_length:
                min_length = length
                min_path = path
        print("The minimum path is {0} with a length of {1}".format(min_path, min_length))


    def solve(self):
        self.solve_for_goals()
        self.solve_for_distances()
        self.solve_for_distances(reset=True)


if __name__ == '__main__':
    with open('day24_input.txt') as f:
        data = tuple([line.strip() for line in f.readlines() if line])
        maze = Maze(data)

    # maze = Maze(tuple([
        # "###########",
        # "#0.1.....2#",
        # "#.#######.#",
        # "#4.......3#",
        # "###########",
    # ]))


    maze.solve()
