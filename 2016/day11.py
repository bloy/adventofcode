#!/usr/bin/env python3

import collections
import itertools


class State(object):
    __slots__ = ('elevator', 'floors')
    def __init__(self, elevator_floor=0,
                 floors=[['AM', 'AG'],
                         ['BG', 'CG', 'DG', 'EG'],
                         ['BM', 'CM', 'DM', 'EM'],
                         [],
                        ]):
        self.elevator = elevator_floor
        self.floors = tuple([frozenset(floor) for floor in floors])

    def __str__(self):
        out = ""
        for (i, floor) in reversed(list(enumerate(self.floors))):
            elevator = " E " if self.elevator == i else "   "
            out += "{0}{1}{2}\n".format(i+1, elevator, " ".join(floor))
        return out

    def __repr__(self):
        return "State(elevator_floor={0}, floors={1})".format(
            self.elevator,
            [list(floor) for floor in self.floors])


    def __eq__(self, other):
        return (
            type(self) == type(other) and
            self.floor_vectors() == other.floor_vectors())

    def floor_vectors(self):
        vectors = []
        for (i, floor) in enumerate(self.floors):
            vectors += [
                (len([thing for thing in floor if thing[-1] == 'M']),
                 len([thing for thing in floor if thing[-1] == 'G']),
                 1 if i == self.elevator else 0)]
        return tuple(vectors)

    def __hash__(self):
        return hash(self.floor_vectors())


    def is_valid(self):
        for floor in self.floors:
            chips = [thing for thing in floor if thing[-1] == 'M']
            gens = [thing for thing in floor if thing[-1] == 'G']
            for chip in chips:
                safe_gen = chip[:-1] + 'G'
                if gens and not safe_gen in gens:
                    return False
        return True

    def is_done(self):
        return (self.floors[3]
                and not self.floors[2]
                and not self.floors[1]
                and not self.floors[0])

    def next_moves(self):
        if not self.is_done():
            nextfloors = [self.elevator + 1, self.elevator - 1]
            if self.elevator == 0: nextfloors = [1]
            if self.elevator == 3: nextfloors = [2]
            contents = self.floors[self.elevator]
            possible_carries = itertools.chain(
                itertools.combinations(contents, 2),
                itertools.combinations(contents, 1)
            )
            moves = itertools.product(nextfloors, possible_carries)
            for move in moves:
                floors = list(self.floors)
                elevator = self.elevator
                next_elevator = move[0]

                floors[elevator] = floors[elevator].difference(move[1])
                floors[next_elevator] = floors[next_elevator].union(move[1])
                newstate = State(
                    elevator_floor=next_elevator,
                    floors=floors)
                if newstate.is_valid():
                    yield newstate


Step = collections.namedtuple('Step', ['length', 'path'])


def find_solution_path(init_state):
    seen = set()

    min_length = None
    min_path = None
    state_queue = collections.deque()
    state_queue.append(Step(0, (init_state, )))

    while state_queue:
        step = state_queue.popleft()
        state = step.path[-1]
        if min_length is None or step.length < min_length:
            if state.is_done():
                min_length = step.length
                min_path = step.path
            elif state not in seen:
                seen.add(state)
                for move in state.next_moves():
                    if move not in seen:
                        state_queue.append(
                            Step(step.length + 1, step.path + (move,))
                        )
    return min_path


def solve1(init_state):
    path = find_solution_path(init_state)
    # for p in path:
        # print(p)
    return len(path)-1

def solve2(init_state):
    path = find_solution_path(init_state)
    for p in path:
        print(p)
    return len(path)-1


if __name__ == '__main__':
    data = [
        ['AG', 'AM'],
        ['BG', 'CG', 'DG', 'EG'],
        ['BM', 'CM', 'DM', 'EM'],
        [],
    ]
    part1_state = State(0, data)

    part2_data = [
        ['AG', 'AM', 'FG', 'FM', 'GG', 'GM'],
        ['BG', 'CG', 'DG', 'EG'],
        ['BM', 'CM', 'DM', 'EM'],
        [],
    ]
    part2_state = State(0, part2_data)

    print(solve1(part1_state))
    print(solve2(part2_state))


