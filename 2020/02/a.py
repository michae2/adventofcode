#!/usr/bin/env python3

import re
import sys

valid = 0
regex = re.compile(r'(\d+)-(\d+) (.): (.*)')
for line in sys.stdin:
    m = regex.match(line)
    lo, hi = map(int, m.group(1, 2))
    ch, pw = m.group(3, 4)
    n = 0
    for pwc in pw:
        if pwc == ch:
            n += 1
    if lo <= n and n <= hi:
        valid += 1
print(valid)
