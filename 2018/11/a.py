#!/usr/bin/env python3

serial_no = int(input())
SIZE = 300

grid = [((x + 10) * y + serial_no) * (x + 10) // 100 % 10 - 5
        for y in range(1, SIZE + 1)
        for x in range(1, SIZE + 1)]

# Memoize partial sums of three horizontal cells.
for y in range(SIZE):
    for x in range(SIZE - 3):
        i = y * SIZE + x
        grid[i] += grid[i + 1] + grid[i + 2]

max_power = grid[0]
max_x = 1
max_y = 1
for y in range(SIZE - 3):
    for x in range(SIZE - 3):
        i = y * SIZE + x
        power = grid[i] + grid[i + SIZE] + grid[i + SIZE * 2]
        if power > max_power:
            max_power = power
            max_x = x + 1
            max_y = y + 1
print(max_x, max_y, sep=',')
