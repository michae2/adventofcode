#!/usr/bin/env python3

serial_no = int(input())
SIZE = 300

grid = [((x + 10) * y + serial_no) * (x + 10) // 100 % 10 - 5
        for y in range(1, SIZE + 1)
        for x in range(1, SIZE + 1)]

# Memoize partial sums of horizontal and vertical runs of cells.  If we iterate
# backward we only need to keep one partial sum per starting cell and direction.
xsums = [0] * SIZE * SIZE
ysums = [0] * SIZE * SIZE

max_power = grid[0]
max_x = 1
max_y = 1
max_d = 1
for y in range(SIZE - 1, -1, -1):
    for x in range(SIZE - 1, -1, -1):
        power = 0
        for d in range(1, min(SIZE - x, SIZE - y)):
            x2 = x + d - 1
            y2 = y + d - 1
            xi = y2 * SIZE + x
            yi = y * SIZE + x2
            ci = y2 * SIZE + x2
            power += xsums[xi]
            power += ysums[yi]
            corner = grid[ci]
            power += corner
            xsums[xi] += corner
            ysums[yi] += corner
            if power > max_power:
                max_power = power
                max_x = x + 1
                max_y = y + 1
                max_d = d
print(max_x, max_y, max_d, sep=',')
