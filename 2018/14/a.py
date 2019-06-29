#!/usr/bin/env python3

goal = int(input())
scoreboard = [3, 7]
elves = [0, 1]
PRINTN = 10
while len(scoreboard) < goal + PRINTN:
    recipe = sum(scoreboard[elf] for elf in elves)
    scoreboard += map(int, str(recipe))
    elves = [(elf + scoreboard[elf] + 1) % len(scoreboard) for elf in elves]
print(*scoreboard[goal:goal + PRINTN], sep='')
