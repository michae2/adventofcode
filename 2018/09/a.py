#!/usr/bin/env python3

import collections
import re

pattern = r'(\d+) players; last marble is worth (\d+) points'
players, last = tuple(map(int, re.match(pattern, input()).groups()))

scores = collections.Counter()
circle = collections.deque()
for m in range(last + 1):
    if m > 0 and m % 23 == 0:
        circle.rotate(7)
        scores[m % players] += m + circle.popleft()
    else:
        circle.rotate(-2)
        circle.appendleft(m)

print(scores.most_common(1)[0][1])
