#!/usr/bin/env python
from collections import defaultdict

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
