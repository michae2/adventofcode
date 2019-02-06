#!/usr/bin/env python3

import sys

commons = set()
for line in sys.stdin:
    box_id = line[:-1]
    for pos in range(len(box_id)):
        common = (pos, box_id[:pos] + box_id[pos + 1:])
        if common in commons:
            print(common[1])
            sys.exit()
        commons.add(common)
