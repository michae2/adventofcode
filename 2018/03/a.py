#!/usr/bin/env python3

import collections
import re
import sys

regex = re.compile(r'#(\d+) @ (\d+),(\d+): (\d+)x(\d+)')
Claim = collections.namedtuple('Claim', 'id x y w h')
claims = [Claim(*map(int, m.group(1, 2, 3, 4, 5)))
          for m in filter(bool, map(regex.match, sys.stdin))]

# Partition the fabric into contiguous "cells" which are rectangles falling
# either completely inside or completely outside each claim.  In the worst case
# there could be as many cells as there are square inches of fabric, but
# hopefully we won't encounter the worst case.

xs = sorted({x for claim in claims for x in (claim.x, claim.x + claim.w)})
ys = sorted({y for claim in claims for y in (claim.y, claim.y + claim.h)})
js = {xs[j]: j for j in range(len(xs))}
ks = {ys[k]: k for k in range(len(ys))}

cells = [[0] * len(ys) for j in range(len(xs))]
overlap_area = 0
for claim in claims:
    for j in range(js[claim.x], js[claim.x + claim.w]):
        for k in range(ks[claim.y], ks[claim.y + claim.h]):
            cells[j][k] += 1
            if cells[j][k] == 2:
                cell_area = (xs[j + 1] - xs[j]) * (ys[k + 1] - ys[k])
                overlap_area += cell_area
print(overlap_area)
