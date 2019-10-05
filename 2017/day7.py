#!env python

import aoc
import collections
import pprint


Program = collections.namedtuple('Program', ['name', 'weight', 'children'])

def read_rows(rows):
    data = []
    for row in rows:
        row = row.split(' -> ')
        children = tuple(row[1].split(', ')) if len(row) == 2 else tuple()
        name, weight = row[0].split(' ')
        weight = int(weight[1:-1])
        program = Program(name, weight, children)
        data.append(program)
    return data


def solve1(data):
    names = set()
    withparents = set()
    for row in data:
        names.add(row.name)
        withparents.update(set(row.children))
    return list(names - withparents)[0]


weights = {}

def total_weight(root, pile):
    if root not in weights:
        child_weights = tuple(total_weight(child, pile)
                              for child in pile[root].children)
        weights[root] = (
            (pile[root].weight + sum(child_weights)),
            pile[root].weight, child_weights)
    return weights[root][0]


def solve2(data):
    root = solve1(data)
    pile = {row.name: row for row in data}
    total_weight(root, pile)
    unbalanced = [(name,
                   pile[name].children,
                   tuple(total_weight(child, pile)
                         for child in pile[name].children))
                  for name in weights if len(set(weights[name][2])) > 1]
    unbalanced_names = set([un[0] for un in unbalanced])
    highest_unbalanced = [test for test in unbalanced
                          if set(test[1]).intersection(unbalanced_names) == set()][0]
    unbalanced_children = [
        (name, pile[name].weight, total_weight(name, pile))
        for name in highest_unbalanced[1]
    ]
    # this just prints out the unbalanced childrein, their total weight
    # and their individual weight. Math is left as an excersize for the reader
    pprint.pprint(unbalanced_children)


rows = """pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)""".split("\n")

if __name__ == '__main__':

    rows = [row for row in aoc.input_lines(day=7)]
    data = read_rows(rows)
    print(solve1(data))
    solve2(data)
