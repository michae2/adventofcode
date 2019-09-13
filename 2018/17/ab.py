#!/usr/bin/env python3

import collections
import enum
import re
import sys

regex = re.compile(r'([xy])=(\d+), [xy]=(\d+)\.\.(\d+)')
Tile = collections.namedtuple('Tile', 'x y')
Soil = enum.Enum('Soil', 'SAND CLAY WET WATER')
syms = ['', '.', '#', '|', '~']

ground = {}
for line in sys.stdin:
    m = regex.match(line)
    z, a, b, c = m.group(1, 2, 3, 4)
    for d in range(int(b), int(c) + 1):
        pos = Tile(int(a), d) if z == 'x' else Tile(d, int(a))
        ground[pos] = Soil.CLAY

min_y = min(t.y for t in ground.keys())
max_y = max(t.y for t in ground.keys())
assert min_y > 0

def left(pos):
    return Tile(pos.x - 1, pos.y)
def right(pos):
    return Tile(pos.x + 1, pos.y)
def above(pos):
    return Tile(pos.x, pos.y - 1)
def below(pos):
    return Tile(pos.x, pos.y + 1)
def scan(pos):
    return ground.get(pos, Soil.SAND)
def blocked(pos):
    return scan(pos) in (Soil.CLAY, Soil.WATER)

PAD = 1
def draw(ground):
    min_x = min(t.x for t in ground.keys()
                if ground[t] in (Soil.WET, Soil.WATER))
    max_x = max(t.x for t in ground.keys()
                if ground[t] in (Soil.WET, Soil.WATER))
    chars = []
    for y in range(min_y, max_y + 1):
        for x in range(min_x - PAD, max_x + 1 + PAD):
            chars.append(syms[scan(Tile(x, y)).value])
        chars.append('\n')
    return ''.join(chars)

# Perform a depth-first search of all the tiles water can reach.
SPRING_X = 500
trickles = [above(Tile(SPRING_X, min_y))]
while trickles:
    # First trickle downward until we hit something.
    trickle = trickles.pop()
    if blocked(trickle):
        continue
    egress = below(trickle)
    while egress not in ground and egress.y <= max_y:
        ground[egress] = Soil.WET
        egress = below(egress)
    if egress.y > max_y or ground[egress] == Soil.WET:
        continue
    # We've hit clay or water.  Trickle leftward.
    trickle = above(egress)
    egress = left(trickle)
    while not blocked(egress) and blocked(below(egress)):
        ground[egress] = Soil.WET
        egress = left(egress)
    if blocked(egress):
        outlet = False
    else:
        outlet = True
        ground[egress] = Soil.WET
        trickles.append(egress)
    # Now trickle rightward.
    egress = right(trickle)
    while not blocked(egress) and blocked(below(egress)):
        ground[egress] = Soil.WET
        egress = right(egress)
    if not blocked(egress):
        outlet = True
        ground[egress] = Soil.WET
        trickles.append(egress)
    # Check whether we found an outlet for this level.
    if not outlet:
        water = left(egress)
        while ground[water] == Soil.WET:
            ground[water] = Soil.WATER
            if scan(above(water)) == Soil.WET:
                trickles.append(above(water))
            water = left(water)

#sys.stdout.write(draw(ground))
counts = collections.Counter(ground.values())
print(counts[Soil.WET] + counts[Soil.WATER])
print(counts[Soil.WATER])
