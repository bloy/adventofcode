#!/usr/bin/python
from itertools import islice
import string
import re

PASSWORD_LENGTH = 8
DOUBLEDOUBLE_RE = re.compile(r'(\w)\1\w*(\w)\2')
TRIPLETS = ["".join(chars)
            for chars in zip(string.ascii_lowercase,
                             string.ascii_lowercase[1:],
                             string.ascii_lowercase[2:])]


def valid_password(password):
    if 'i' in password or 'l' in password or 'o' in password:
        return False
    if not DOUBLEDOUBLE_RE.search(password):
        return False
    if not any(((triple in password) for triple in TRIPLETS)):
        return False
    return True


def passwords(current):
    most_recent = list(current)
    while most_recent != list('z' * PASSWORD_LENGTH):
        for i in range(-1, -1 * PASSWORD_LENGTH, -1):
            if most_recent[i] == 'z':
                most_recent[i] = 'a'
            else:
                most_recent[i] = chr(ord(most_recent[i])+1)
                break
        pwd = "".join(most_recent)
        if valid_password(pwd):
            yield pwd


if __name__ == '__main__':
    current_password = 'hepxcrrq'
    for index, word in enumerate(islice(passwords(current_password), 10)):
        print(index, ':', word)
