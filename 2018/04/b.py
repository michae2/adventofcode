#!/usr/bin/env python3

import collections
import datetime
import itertools
import operator
import sys

records = [(datetime.datetime.strptime(r[0], '[%Y-%m-%d %H:%M'), r[1])
           for r in map(lambda line: line.rstrip().split(sep='] '), sys.stdin)]

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

naps = [(guard, list(itertools.accumulate(deltas)))
        for guard, deltas in nap_deltas.items()]

modes = [(guard, max(enumerate(minutes), key=operator.itemgetter(1)))
         for guard, minutes in naps]

guard, (mode, count) = max(modes, key=lambda g: g[1][1])
print (guard * mode)
