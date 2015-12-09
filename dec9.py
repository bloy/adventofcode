#!/usr/bin/env python
import re


class PathFinder(object):
    REGEX = re.compile(r'^(.*)\sto\s(.*)\s=\s(\d+)$')

    def __init__(self, lines):
        self.places = {}
        for line in lines:
            (place1, place2, distance) = self.REGEX.match(line.strip()).groups()
            distance = int(distance)
            self.add_distance(place1, place2, distance)
            self.add_distance(place2, place1, distance)
        self.place_set = set(self.places.keys())

    def add_distance(self, place_a, place_b, distance):
        if place_a not in self.places:
            self.places[place_a] = {}
        self.places[place_a][place_b] = distance

    def find_shortest_helper(self, current_place, current_distance, visited_list):
        possible_set = self.place_set.difference(set(visited_list))
        if len(possible_set) == 0:
            return (current_distance, visited_list)
        shortest = None
        shortest_distance = 65536
        for place in possible_set:
            distance = self.places[current_place][place]
            new_visited_list = visited_list + [place]
            path = self.find_shortest_helper(place,
                                             current_distance + distance,
                                             new_visited_list)
            if path[0] < shortest_distance:
                shortest_distance = path[0]
                shortest = path
        return shortest


    def find_shortest(self):
        shortest_distance = 65536
        shortest = None
        for place in self.place_set:
            path = self.find_shortest_helper(place, 0, [place])
            if path[0] < shortest_distance:
                shortest_distance = path[0]
                shortest = path

        return shortest


if __name__ == '__main__':
    # lines = [
        # 'London to Dublin = 464\n',
        # 'London to Belfast = 518\n',
        # 'Dublin to Belfast = 141\n',
    # ]
    with open('input/day_9') as lines:
        finder = PathFinder(lines)

    print(finder.find_shortest())
