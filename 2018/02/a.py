#!/usr/bin/env python3

import sys

twos = 0
threes = 0
for line in sys.stdin:
    box_id = line[:-1]
    counts = {}
    for letter in box_id:
        counts[letter] = counts.get(letter, 0) + 1
    if 2 in counts.values():
        twos += 1
    if 3 in counts.values():
        threes += 1
print(twos * threes)
