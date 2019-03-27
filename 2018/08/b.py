#!/usr/bin/env python3

import itertools
import sys

numbers = map(int, sys.stdin.read(None).split())

def node_value():
    c = next(numbers)
    m = next(numbers)
    kids = [node_value() for _ in range(c)]
    return sum(e if not c else 0 if e < 1 or e > c else kids[e - 1]
               for e in itertools.islice(numbers, m))

print(node_value())
