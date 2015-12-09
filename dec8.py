#!/usr/bin/env python
import re

ESCAPE_RE = re.compile(r'''\\(x[0-9a-fA-F]{2}|\\|"|')''')
ENCODE_RE = re.compile(r'''(\\|\")''')

def totals_from_lines(lines):
    file_total = 0
    string_total = 0
    encoded_total = 0

    for line in lines:
        line = line.strip()
        string = ESCAPE_RE.sub('5', line)[1:-1]
        encoded = '"' + ENCODE_RE.sub(lambda match: "\\" + match.group(0), line) + '"'
        print(line, encoded)
        file_total += len(line)
        string_total += len(string)
        encoded_total += len(encoded)

    return (file_total, string_total, encoded_total)

if __name__ == '__main__':
    with open('input/day_8') as lines:
        (file_total, string_total, encoded_total) = totals_from_lines(lines)

    print('file_total:', file_total)
    print('string_total:', string_total)
    print('encoded total:', encoded_total)
    print('difference:', file_total - string_total)
    print('encoded difference:', encoded_total - file_total)
