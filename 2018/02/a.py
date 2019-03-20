#!/usr/bin/env python3

import collections
import sys

twos = 0
threes = 0
for line in sys.stdin:
    box_id = line.rstrip()
    counts = collections.Counter(box_id)
    if 2 in counts.values():
        twos += 1
    if 3 in counts.values():
        threes += 1
print(twos * threes)
