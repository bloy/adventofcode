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
        Reindeer('Comet', 14, 10, 127),
        Reindeer('Dancer', 16, 11, 162),
    )
    for seconds in (959, 963, 1000):
        for reindeer in data:
            print("reindeer:", reindeer.name,
                  "time:", seconds,
                  "distance:", get_distance(reindeer, seconds))
