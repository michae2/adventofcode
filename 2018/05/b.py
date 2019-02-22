#!/usr/bin/env python3

def react(a, b):
    return a.lower() == b.lower() and a != b

reactant = input()
product = [' '] * len(reactant)
minp = len(reactant)
utypes = set(reactant.lower())
for utype in utypes:
    p = 0
    for unit in reactant:
        if unit.lower() == utype:
            continue
        elif p > 0 and react(unit, product[p - 1]):
            p -= 1
        else:
            product[p] = unit
            p += 1
    minp = min(p, minp)
print(minp)
