#!/usr/bin/env python
from itertools import count, combinations_with_replacement
import pprint

SPELLS = {'Magic Missile', 'Drain', 'Shield', 'Poison', 'Recharge'}
SPELL_MANA = {
    'Magic Missile': 53,
    'Drain': 73,
    'Shield': 113,
    'Poison': 173,
    'Recharge': 229
}
SPELL_TIMERS = {
    'Magic Missile': 0,
    'Drain': 0,
    'Shield': 6,
    'Poison': 6,
    'Recharge': 5
}
BOSS_START_HP = 71
BOSS_DAMAGE = 10
HERO_START_HP = 50
HERO_START_MANA = 500


def spell_effects(effects, boss_hp, hero_mana):
    new_effects = dict()
    for effect in effects:
        if effect == 'Poison':
            boss_hp -= 3
        elif effect == 'Recharge':
            hero_mana += 101
        if effects[effect] < 0:
            new_effects[effect] = effects[effect] - 1
    return new_effects, boss_hp, hero_mana


def battle_for_spell_order(spell_order):
    boss_hp = BOSS_START_HP
    hero_hp = HERO_START_HP
    hero_mana = HERO_START_MANA
    hero_armor = 0
    spent_mana = 0
    effects = dict()

    for spell in spell_order:
        if hero_hp <= 0:
            return -100
        (effects, boss_hp, hero_mana) = spell_effects(
            effects, boss_hp, hero_mana)
        if boss_hp <= 0:
            return spent_mana
        if SPELL_MANA[spell] > hero_mana:
            return -100
        if spell in effects:
            return -100
        if spell == 'Magic Missle':
            boss_hp -= 4
        elif spell == 'Drain':
            boss_hp -= 2
            hero_hp += 2
        else:
            effects[spell] = SPELL_TIMERS[spell]
        spent_mana += SPELL_MANA[spell]
        hero_mana -= SPELL_MANA[spell]

        if 'Shield' in effects:
            hero_armor = 7
        else:
            hero_armor = 0
        (effects, boss_hp, hero_mana) = spell_effects(
            effects, boss_hp, hero_mana)
        if boss_hp <= 0:
            return spent_mana
        hero_hp -= max(1, BOSS_DAMAGE - hero_armor)
        if hero_hp <= 0:
            return -100
    return -1




if __name__ == '__main__':
    best_list = None
    best_mana = None
    for num_spells in count():
        for spell_list in combinations_with_replacement(SPELLS, num_spells):
            result = battle_for_spell_order(spell_list)
            if result >= 0:
                if result < best_mana:
                    best_mana = result
                    best_list = spell_list
        if best_list:
            break
        pprint.pprint(num_spells)
    pprint.pprint((best_mana, best_list))
