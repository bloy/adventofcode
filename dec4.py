#!/usr/bin/env python

import hashlib

secret = 'iwrupvqb'
number = -1
hash_result = 'xxxxxxxxxxxxxxxxxxx'
while hash_result[:6] != '000000':
    number = number + 1
    secret_string = secret + str(number)
    hash_result = hashlib.md5(secret_string.encode('utf-8')).hexdigest()

print(number)
