#!/usr/bin/env python3

import collections
import sys

coords = [tuple(map(int, line.rstrip().split(', ')))
          for line in sys.stdin]

assert len(coords) > 0

# Areas that extend beyond the farthest coordinate in any direction will be
# infinite, so we don't need to explore beyond the outermost coordinates.
xs, ys = zip(*coords)
min_x = min(xs)
max_x = max(xs)
min_y = min(ys)
max_y = max(ys)

# Visit every location by "growing" outward from each coord one step at a time.
areas = collections.Counter()
unvisited = {(x, y)
             for x in range(min_x, max_x + 1)
             for y in range(min_y, max_y + 1)}
visits = {}

for c, (x, y) in enumerate(coords, start=1):
    areas[c] = 0
    if (x, y) in unvisited:
        visits[(x, y)] = c
        unvisited.remove((x, y))
    elif (x, y) in visits:
        # It's a tie from the beginning! (This seems unlikely...)
        visits[(x, y)] = 0

while len(visits):
    tovisit = {}
    for (x, y), c in visits.items():
        if c in areas:
            areas[c] += 1
        for x1, y1 in [(x - 1, y), (x + 1, y), (x, y - 1), (x, y + 1)]:
            if (x1, y1) in unvisited:
                tovisit[(x1, y1)] = c
                unvisited.remove((x1, y1))
            elif (x1, y1) in tovisit:
                if tovisit[(x1, y1)] != c:
                    # It's a tie!
                    tovisit[(x1, y1)] = 0
            elif c in areas and (x1 < min_x or x1 > max_x or
                                 y1 < min_y or y1 > max_y):
                # Area is infinite.
                del areas[c]
    visits = tovisit

print(areas.most_common(1)[0][1])
