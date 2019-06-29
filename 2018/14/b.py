#!/usr/bin/env python3

goal = list(map(int, input()))
scoreboard = [3, 7]
elf1 = 0
elf2 = 1

while True:
    recipe = scoreboard[elf1] + scoreboard[elf2]
    if recipe >= 10:
        scoreboard.append(1)
        if scoreboard[-len(goal):] == goal:
            break
    scoreboard.append(recipe % 10)
    if scoreboard[-len(goal):] == goal:
        break
    elf1 = (elf1 + scoreboard[elf1] + 1) % len(scoreboard)
    elf2 = (elf2 + scoreboard[elf2] + 1) % len(scoreboard)

print(len(scoreboard) - len(goal))
