#!/usr/bin/env python3

import functools
import sys

initial = input()[15:]
input()
rules = {line[0:5]: line[9:10] for line in sys.stdin}

assert rules.get('.....', '.') == '.'
C = 2
GENS = 50000000000
horizon = C * (GENS + 1)

def next_gen(pots):
    children = ''
    for p in range(C, len(pots) - C):
        children += rules.get(pots[p - C:p + C + 1], '.')
    return children

# Use a Hashlife-like algorithm.

class Pot:
    def __init__(self, plants):
        self.level = 0
        self.size = 1
        self.num = plants
        self.pos_sum = plants
        self.neg_sum = -plants
        self.max_step = 2 ** (-C - 1)
    def __str__(self):
        return '#' if self.num > 0 else '.'

dead = Pot(0)
live = Pot(1)

def build(potstr):
    return [live if pot == '#' else dead for pot in potstr]

def build_level(pots):
    return [Pots(pots[i], pots[i + 1]) for i in range(0, len(pots), 2)]

@functools.lru_cache(maxsize=None)
class Pots:
    def __init__(self, l, r):
        assert l.level == r.level
        assert l.size == r.size
        self.l = l
        self.r = r
        self.level = l.level + 1
        self.size = l.size + r.size
        self.num = l.num + r.num
        self.pos_sum = l.pos_sum + r.pos_sum + r.num * l.size
        self.neg_sum = r.neg_sum + l.neg_sum - l.num * r.size
        self.max_step = 2 ** (self.level - C - 1)
    def __str__(self):
        return str(self.l) + str(self.r)

    @functools.lru_cache(maxsize=None)
    def step(self, gens):
        assert gens <= self.max_step
        if self.max_step == 1 and gens == 1:
            children = next_gen(str(self))
            pots = build(children)
            while len(pots) > 1:
                pots = build_level(pots)
            return pots[0]
        c = Pots(self.l.r, self.r.l)
        if gens == 0:
            return c
        g1 = min(gens, c.max_step)
        l1 = self.l.step(g1)
        r1 = self.r.step(g1)
        c1 = c.step(g1)
        g2 = gens - g1
        l2 = Pots(l1, c1).step(g2)
        r2 = Pots(c1, r1).step(g2)
        return Pots(l2, r2)

pos = build(initial)
neg = build('.' * len(initial))
empty = dead
while len(pos) > 1 or pos[0].size < 2 * (len(initial) + horizon):
    if len(pos) % 2 == 1:
        pos.append(empty)
        neg.append(empty)
    pos = build_level(pos)
    neg = build_level(neg)
    empty = Pots(empty, empty)

universe = Pots(neg[0], pos[0])
universe = universe.step(GENS)
print(universe.l.neg_sum + universe.r.pos_sum - universe.r.num)
