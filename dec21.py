#!/usr/bin/env python
import collections
import itertools
import pprint
import math

Equipment = collections.namedtuple('Equipment',
                                   ['name', 'cost', 'damage', 'armor'])
Character = collections.namedtuple('Character',
                                   ['name', 'hitpoints', 'damage', 'armor'])
WEAPONS = (
    Equipment("Dagger", 8, 4, 0),
    Equipment("Shortsword", 10, 5, 0),
    Equipment("Warhammer", 25, 6, 0),
    Equipment("Longsword", 40, 7, 0),
    Equipment("Greataxe", 74, 8, 0),
)

ARMOR = (
    Equipment("Clothing", 0, 0, 0),
    Equipment("Leather", 13, 0, 1),
    Equipment("Chainmail", 31, 0, 2),
    Equipment("Splintmail", 53, 0, 3),
    Equipment("Bandedmail", 75, 0, 4),
    Equipment("Platemail", 102, 0, 5),
)

RINGS = (
    Equipment("No ring", 0, 0, 0),
    Equipment("Damage +1", 25, 1, 0),
    Equipment("Damage +2", 50, 2, 0),
    Equipment("Damage +3", 100, 3, 0),
    Equipment("Defense +1", 20, 0, 1),
    Equipment("Defense +2", 40, 0, 2),
    Equipment("Defense +3", 80, 0, 3),
)

def attacker_turns(attacker, defender):
    return math.ceil(defender.hitpoints /
                     max(attacker.damage - defender.armor, 1))

def hero_wins(hero, boss):
    hero_turns = attacker_turns(hero, boss)
    boss_turns = attacker_turns(boss, hero)
    if hero_turns <= boss_turns:
        return True
    else:
        return False


if __name__ == '__main__':
    boss = Character('Boss', 109, 8, 2)

    cheapest_cost = 65535
    cheapest_set = None

    equipment_sets = (equipment_set for equipment_set
                      in itertools.product(WEAPONS, ARMOR, RINGS, RINGS)
                      if equipment_set[2].cost == 0 or equipment_set[2] != equipment_set[3])

    for equipment_set in equipment_sets:
        cost = sum(e.cost for e in equipment_set)
        damage = sum(e.damage for e in equipment_set)
        armor = sum(e.armor for e in equipment_set)
        hero = Character('Hero', 100, damage, armor)
        hero_turns = attacker_turns(hero, boss)
        boss_turns = attacker_turns(boss, hero)
        if cost < cheapest_cost and hero_turns <= boss_turns:
            cheapest_cost = cost
            cheapest_set = equipment_set
    pprint.pprint(cheapest_cost)
    pprint.pprint(cheapest_set)

