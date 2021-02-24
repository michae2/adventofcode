#!/usr/bin/env python3

import sys

es = set()
for line in sys.stdin:
    e = int(line)
    d = 2020 - e
    if d in es:
        print(d * e)
        break
    es.add(e)
