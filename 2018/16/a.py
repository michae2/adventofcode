#!/usr/bin/env python3

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

ambiguous = 0

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

    matches = 0
    for op in ops:
        reg = reg1.copy()
        op(reg, a, b, c)
        if reg == reg2:
            matches += 1
            if matches >= 3:
                ambiguous += 1
                break

    before = input()

print(ambiguous)
