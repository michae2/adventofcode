#!/usr/bin/env python3

import sys

coords = [tuple(map(int, line.rstrip().split(', ')))
          for line in sys.stdin]

assert len(coords) > 0

# Distances will always increase once we're outside of the median square. So
# start there and walk outward, increasing the distance by an increment which
# itself increases when we pass by a coordinate.

xs, ys = zip(*coords)
xs = sorted(xs)
ys = sorted(ys)

m = len(coords) // 2
mx = xs[m]
my = ys[m]

start_dist = 0
for x, y in coords:
    start_dist += abs(mx - x)
    start_dist += abs(my - y)

left = [mx - x for x in reversed(xs[:m])]
right = [x - mx for x in xs[m:]]
up = [my - y for y in reversed(ys[:m])]
down = [y - my for y in ys[m:]]

def incr(zs):
    d = len(coords) % 2
    x = 0
    for z in zs:
        while x < z:
            yield d
            x += 1
        d += 2
    while True:
        yield d

def walk(zs, dist, first):
    d = incr(zs)
    if not first:
        dist += next(d)
    while dist < 10000:
        yield dist
        dist += next(d)

area = 0
for dist in walk(left, start_dist, False):
    area += len(list(walk(up, dist, False)))
    area += len(list(walk(down, dist, True)))
for dist in walk(right, start_dist, True):
    area += len(list(walk(up, dist, False)))
    area += len(list(walk(down, dist, True)))
print(area)
