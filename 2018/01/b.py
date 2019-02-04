#!/usr/bin/env python3

import itertools
import sys

changes = map(int, sys.stdin)
changes = itertools.cycle(changes)
changes = iter(changes)
freqs = set()
freq = 0
while freq not in freqs:
    freqs.add(freq)
    freq += next(changes)
print(freq)
