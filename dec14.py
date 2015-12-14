#!/usr/bin/env python
import collections

Reindeer = collections.namedtuple('Reindeer', ['name', 'speed',
                                               'fly_time', 'rest_time'])

def get_distance(reindeer, seconds):
    speed = reindeer.speed
    fly_time = reindeer.fly_time
    rest_time = reindeer.rest_time
    cycle_time = fly_time + rest_time
    distance = (fly_time * (seconds / cycle_time) +
                min(fly_time, seconds % cycle_time)) * speed
    return distance

def part1(data):
    best = 0
    for reindeer in data:
        distance = get_distance(reindeer, 2503)
        if best < distance:
            best = distance
        print("reindeer:", reindeer.name,
              "distance:", distance)
    print best

def part2(data):
    points = collections.Counter()
    for second in range(2503):
        leader = None
        leader_distance = 0
        for deer in data:
            distance = get_distance(deer, second+1)
            if leader_distance < distance:
                leader_distance = distance
                leader = deer.name
        points[leader] += 1
    print(points)


if __name__ == '__main__':
    data = (
        Reindeer('Rudolph', 22, 8, 165),
        Reindeer('Cupid', 8, 17, 114),
        Reindeer('Prancer', 18, 6, 103),
        Reindeer('Donner', 25, 6, 145),
        Reindeer('Dasher', 11, 12, 125),
        Reindeer('Comet', 21, 6, 121),
        Reindeer('Blitzen', 18, 3, 50),
        Reindeer('Vixen', 20, 4, 75),
        Reindeer('Dancer', 7, 20, 119),
    )
    part1(data)
    part2(data)
