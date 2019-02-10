#!/usr/bin/env python3

import collections
import re
import sys

regex = re.compile(r'#(\d+) @ (\d+),(\d+): (\d+)x(\d+)')
Claim = collections.namedtuple('Claim', 'id x y w h')
claims = [Claim(*map(int, m.group(1, 2, 3, 4, 5)))
          for m in filter(bool, map(regex.match, sys.stdin))]
claims = sorted(claims, key=lambda claim: claim.x)
overlapping = [False] * len(claims)
for i, claim in enumerate(claims):
    for j in range(i + 1, len(claims)):
        claim2 = claims[j]
        if claim2.x >= claim.x + claim.w:
            break
        elif overlapping[i] and overlapping[j]:
            continue
        elif claim.y < claim2.y + claim2.h and claim2.y < claim.y + claim.h:
            overlapping[i] = True
            overlapping[j] = True
    if not overlapping[i]:
        print(claim.id)
        sys.exit()
