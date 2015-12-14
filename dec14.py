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
    for seconds in (2503,):
        for reindeer in data:
            print("reindeer:", reindeer.name,
                  "time:", seconds,
                  "distance:", get_distance(reindeer, seconds))
