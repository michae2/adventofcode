#!/usr/bin/env python3

import collections
import sys

WORKERS = 5
SECONDS = 60

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

idle = WORKERS
busy = {}
time = 0

while len(ready) or idle < WORKERS:
    while len(ready) and idle:
        step = min(ready)
        ready.remove(step)
        busy[step] = SECONDS + ord(step) - ord('A') + 1
        idle -= 1
    interval = min(busy.values())
    time += interval
    for step in busy:
        busy[step] -= interval
    done = [step for step, seconds in busy.items() if seconds == 0]
    for step in done:
        del busy[step]
        idle += 1
        for requiree in required[step]:
            requires[requiree].remove(step)
            if len(requires[requiree]) == 0:
                ready.add(requiree)
print(time)
