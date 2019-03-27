#!/usr/bin/env python3

import itertools
import sys

numbers = map(int, sys.stdin.read(None).split())

def node_sum():
    c = next(numbers)
    m = next(numbers)
    return sum(node_sum() for _ in range(c)) + sum(itertools.islice(numbers, m))

print(node_sum())
