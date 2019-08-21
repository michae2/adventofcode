#!/usr/bin/env python3

import sys

def addr(reg, a, b, c):
    reg[c] = reg[a] + reg[b]

def addi(reg, a, b, c):
    reg[c] = reg[a] + b

def mulr(reg, a, b, c):
    reg[c] = reg[a] * reg[b]

def muli(reg, a, b, c):
    reg[c] = reg[a] * b

def banr(reg, a, b, c):
    reg[c] = reg[a] & reg[b]

def bani(reg, a, b, c):
    reg[c] = reg[a] & b

def borr(reg, a, b, c):
    reg[c] = reg[a] | reg[b]

def bori(reg, a, b, c):
    reg[c] = reg[a] | b

def setr(reg, a, b, c):
    reg[c] = reg[a]

def seti(reg, a, b, c):
    reg[c] = a

def gtir(reg, a, b, c):
    reg[c] = 1 if a > reg[b] else 0

def gtri(reg, a, b, c):
    reg[c] = 1 if reg[a] > b else 0

def gtrr(reg, a, b, c):
    reg[c] = 1 if reg[a] > reg[b] else 0

def eqir(reg, a, b, c):
    reg[c] = 1 if a == reg[b] else 0

def eqri(reg, a, b, c):
    reg[c] = 1 if reg[a] == b else 0

def eqrr(reg, a, b, c):
    reg[c] = 1 if reg[a] == reg[b] else 0

ops = [addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri,
       gtrr, eqir, eqri, eqrr]

ops = [(op, set(range(len(ops)))) for op in ops]

before = input()
while before.startswith("Before"):
    instr = input()
    after = input()
    blank = input()

    assert after.startswith("After")
    assert not blank

    reg1 = eval(before[8:])
    reg2 = eval(after[8:])
    opcode, a, b, c = map(int, instr.split())

    for op, opcodes in ops:
        reg = reg1.copy()
        op(reg, a, b, c)
        if reg != reg2:
            opcodes.discard(opcode)

    before = input()

# Use process of elimination to find the mapping from opcodes to ops.
opmap = {}
found = True
while found:
    found = set()
    for op, opcodes in ops:
        if len(opcodes) == 1:
            opcode = opcodes.pop()
            opmap[opcode] = op
            found.add(opcode)
    for op, opcodes in ops:
        opcodes -= found

assert len(opmap) == len(ops)

blank = input()
assert not blank

# Now run the test program.
reg = [0] * 4

for instr in sys.stdin:
    opcode, a, b, c = map(int, instr.split())
    op = opmap[opcode]
    op(reg, a, b, c)

print(reg[0])
