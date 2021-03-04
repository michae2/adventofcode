#!/usr/bin/env python3

import re
import sys

valid = 0
regex = re.compile(r'(\d+)-(\d+) (.): (.*)')
for line in sys.stdin:
    m = regex.match(line)
    a, b = map(int, m.group(1, 2))
    ch, pw = m.group(3, 4)
    if (a <= len(pw) and pw[a - 1] == ch) ^ (b <= len(pw) and pw[b - 1] == ch):
        valid += 1
print(valid)
