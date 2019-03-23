#!/usr/bin/env python3

import collections
import sys

required = collections.defaultdict(set)
requires = collections.defaultdict(set)
for line in sys.stdin:
    required[line[5:6]].add(line[36:37])
    requires[line[5:6]]
    requires[line[36:37]].add(line[5:6])
    required[line[36:37]]

ready = set()
for step, reqs in requires.items():
    if len(reqs) == 0:
        ready.add(step)

while len(ready):
    step = min(ready)
    ready.remove(step)
    sys.stdout.write(step)
    for requiree in required[step]:
        requires[requiree].remove(step)
        if len(requires[requiree]) == 0:
            ready.add(requiree)
sys.stdout.write("\n")
