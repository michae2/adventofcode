#!/usr/bin/env python3

def react(a, b):
    return a.lower() == b.lower() and a != b

reactant = input()
product = [' '] * len(reactant)
p = 0
for unit in reactant:
    if p > 0 and react(unit, product[p - 1]):
        p -= 1
    else:
        product[p] = unit
        p += 1
print(p)
