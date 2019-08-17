#!/usr/bin/env python3

import operator
import sys

rows = [line[:-1] for line in sys.stdin]
h = len(rows)
w = len(rows[0])
syms = ''.join(rows)

class Troglobyte:
    def __init__(self, pos):
        self.pos = pos

class Wall(Troglobyte):
    def __str__(self):
        return "#"

class Unit(Troglobyte):
    def __init__(self, pos, attack):
        self.attack = attack
        self.hp = 200
        super().__init__(pos)
    def alive(self):
        return self.hp > 0

class Elf(Unit):
    def __str__(self):
        return "E"
    @staticmethod
    def enemy():
        return Goblin

class Goblin(Unit):
    def __init__(self, pos):
        super().__init__(pos, 3)
    def __str__(self):
        return "G"
    @staticmethod
    def enemy():
        return Elf

class ElfKilled(Exception):
    pass

def gather(cavern, clade):
    trogs = (trog for trog in cavern.values() if isinstance(trog, clade))
    return sorted(trogs, key=operator.attrgetter('pos'))

def adjacent(pos):
    if pos // w > 0:
        yield pos - w
    if pos % w > 0:
        yield pos - 1
    if pos % w < w - 1:
        yield pos + 1
    if pos // w < h - 1:
        yield pos + w

def do_attack(cavern, unit, target):
    target.hp -= unit.attack
    if not target.alive():
        if isinstance(target, Elf):
            raise ElfKilled
        del cavern[target.pos]

def find_attack(cavern, unit):
    targets = []
    for pos in adjacent(unit.pos):
        if pos in cavern and isinstance(cavern[pos], unit.enemy()):
            targets.append(cavern[pos])
    if not targets:
        return None
    return min(targets, key=operator.attrgetter('hp'))

def do_move(cavern, unit, pos):
    del cavern[unit.pos]
    cavern[pos] = unit
    unit.pos = pos

def find_move(cavern, unit, targets):
    for pos in adjacent(unit.pos):
        if pos in cavern and isinstance(cavern[pos], unit.enemy()):
            return -1

    attack_range = set()
    for target in targets:
        for pos in adjacent(target.pos):
            if pos not in cavern:
                attack_range.add(pos)

    # Perform a breadth-first search from unit to find the nearest positions
    # in attack range.
    nearest = []
    explore = []
    path = {} # map from pos to initial step
    for pos in adjacent(unit.pos):
        if pos not in cavern:
            explore.append(pos)
            path[pos] = pos

    while explore:
        for pos in explore:
            if pos in attack_range:
                nearest.append(pos)
        if nearest:
            break
        new_edge = []
        for pos in explore:
            for new in adjacent(pos):
                if new not in cavern and new not in path:
                    new_edge.append(new)
                    path[new] = path[pos]
        explore = new_edge

    if not nearest:
        return -1

    dest = min(nearest)
    return path[dest]

def do_turn(cavern, unit):
    targets = gather(cavern, unit.enemy())
    if not targets:
        return False

    pos = find_move(cavern, unit, targets)
    if pos >= 0:
        do_move(cavern, unit, pos)

    target = find_attack(cavern, unit)
    if target:
        do_attack(cavern, unit, target)

    return True

def do_round(cavern):
    for unit in gather(cavern, Unit):
        if not unit.alive():
            continue
        if not do_turn(cavern, unit):
            return False
    return True

def do_battle(syms, elf_attack):
    cavern = {}
    for pos, sym in enumerate(syms):
        if sym == '#':
            cavern[pos] = Wall(pos)
        elif sym == 'E':
            cavern[pos] = Elf(pos, elf_attack)
        elif sym == 'G':
            cavern[pos] = Goblin(pos)

    rounds = 0
    while do_round(cavern):
        rounds += 1

    total_hp = sum(unit.hp for unit in gather(cavern, Unit))
    return rounds * total_hp

elf_attack = 4
while True:
    try:
        print(do_battle(syms, elf_attack))
        #print(elf_attack)
        sys.exit()
    except ElfKilled:
        elf_attack += 1
