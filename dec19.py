#!/usr/bin/env python
from collections import defaultdict
from random import shuffle

def parse_lines(lines):
    transforms = defaultdict(list)
    for line in lines:
        line = line.strip()
        (key, value) = line.split(' => ')
        transforms[key].append(value)
    return transforms

def substitutions(medicine, old, new):
    splits = medicine.split(old)
    return [new.join((old.join(splits[:x]), old.join(splits[x:])))
            for x in range(1, len(splits))]


def possible_transforms(transforms, medicine):
    transformations = set()
    for old in transforms.keys():
        for new in transforms[old]:
            for substitution in substitutions(medicine, old, new):
                transformations.add(substitution)
    return transformations


def find_shortest_path(start, target, transforms):

    replacements = [(''.join(reversed(k)), ''.join(reversed(r)))
                    for k in transforms for r in transforms[k]]
    molecule = ''.join(reversed(target))

    count = 0
    while molecule != start:
        tmp = molecule
        for new, replacement in replacements:
            if replacement in molecule:
                molecule = molecule.replace(replacement, new, 1)
                count += 1
                print(''.join(reversed(molecule)))

        if tmp == molecule:
            count = 0
            molecule = ''.join(reversed(target))
            shuffle(replacements)
    return count


if __name__ == '__main__':
    with open('input/day_19') as lines:
        transforms = parse_lines(lines)
    medicine = (
        'CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiTh'
        'CaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCa'
        'SiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgY'
        'CaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRn'
        'FYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRn'
        'FArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiTh'
        'CaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr'
    )

    print(len(possible_transforms(transforms, medicine)))
    print(find_shortest_path('e', medicine, transforms))
