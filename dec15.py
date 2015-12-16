#!/usr/bin/env python
import collections
import pprint

Ingredient = collections.namedtuple(
    'Ingredient',
    ['name', 'capacity', 'durability', 'flavor', 'texture', 'calories'])


def permutations(limit, ingredients):
    if len(ingredients) == 1:
        yield (limit,)
    else:
        for x in range(limit+1):
            for permutation in permutations(limit-x, ingredients[1:]):
                yield (x,) + permutation

def find_score(permutation, ingredients):
    amounts = [[amount*attribute for attribute in ingredient[1:-1]]
               for (amount, ingredient) in zip(permutation, ingredients)]
    totals = [sum(score_parts) for score_parts in zip(*amounts)]
    totals = [t if t > 0 else 0 for t in totals]
    total = reduce(lambda x,y: x * y, totals)
    return total


def find_optimal_ingredients(limit, ingredients):
    best = 0
    best_permutation = None
    for permutation in permutations(limit, ingredients):
        score = find_score(permutation, ingredients)
        if score > best:
            best_permutation = zip(permutation, ingredients)
            best = score
    return (best, best_permutation)

if __name__ == '__main__':
    ingredients = (
        Ingredient('Frosting', 4, -2, 0, 0, 5),
        Ingredient('Candy', 0, 5, -1, 0, 8),
        Ingredient('Butterscotch', -1, 0, 5, 0, 6),
        Ingredient('Sugar', 0, 0, -2, 2, 1),
    )
    pprint.pprint(find_optimal_ingredients(100, ingredients))
