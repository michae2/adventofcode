#!/usr/bin/env python3

import collections
import datetime
import itertools
import operator
import sys

records = [(datetime.datetime.strptime(r[0], '[%Y-%m-%d %H:%M'), r[1])
           for r in map(lambda line: line.rstrip().split(sep='] '), sys.stdin)]

# Sort by time to assign each event to a guard.
records.sort(key=operator.itemgetter(0))
nap_deltas = collections.defaultdict(lambda: [0] * 60)
guard = None
for time, msg in records:
    if msg.startswith('Guard'):
        guard = int(msg[7:].split()[0])
    elif msg.startswith('falls'):
        nap_deltas[guard][time.minute] += 1
    else:
        nap_deltas[guard][time.minute] -= 1

# Integrate nap-deltas per minute to get naps per minute.
naps = [(guard, list(itertools.accumulate(deltas)))
        for guard, deltas in nap_deltas.items()]
guard, minutes = max(naps, key=lambda g: sum(g[1]))
mode, count = max(enumerate(minutes), key=operator.itemgetter(1))
print(guard * mode)
