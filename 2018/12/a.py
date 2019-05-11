#!/usr/bin/env python3

import sys

initial = input()[15:]
input()
rules = {line[0:5]: line[9:10] for line in sys.stdin}

# The fastest that plants can spread (the "speed of light") is 2 pots per
# generation.
assert rules.get('.....', '.') == '.'
C = 2
GENS = 20
horizon = C * (GENS + 1)

def next_gen(pots):
    children = ''
    for p in range(C, len(pots) - C):
        children += rules.get(pots[p - C:p + C + 1], '.')
    return children

pots = '.' * horizon + initial + '.' * horizon
for _ in range(GENS):
    pots = '.' * C + next_gen(pots) + '.' * C

total = 0
for p, pot in enumerate(pots):
    if pot == '#':
        total += p - horizon
print(total)
