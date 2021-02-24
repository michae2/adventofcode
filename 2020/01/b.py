#!/usr/bin/env python3

import sys

es = list()
tab = {}
for line in sys.stdin:
    e = int(line)
    d = 2020 - e
    if d in tab:
        print(tab[d] * e)
        break
    tab.update({e0 + e: e0 * e for e0 in es})
    es.append(e)
