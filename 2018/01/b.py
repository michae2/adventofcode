#!/usr/bin/env python3

import itertools
import sys

changes = map(int, sys.stdin)
changes = itertools.cycle(changes)
changes = changes.__iter__()
freqs = set()
freq = 0
while freq not in freqs:
    freqs.add(freq)
    freq += changes.__next__()
print(freq)
