#!/usr/bin/env python3

import itertools
import operator
import sys

rows = [line[:-1] for line in sys.stdin]
w = len(rows[0])
tracks = ''.join(rows)

cart_moves = {'>': +1, 'v': +w, '<': -1, '^': -w}

lcurve   = {+1: +w, +w: +1, -1: -w, -w: -1}  # \
rcurve   = {+1: -w, +w: -1, -1: +w, -w: +1}  # /
left     = {+1: -w, +w: +1, -1: +w, -w: -1}  # +
straight = {+1: +1, +w: +w, -1: -1, -w: -w}  # + | -
right    = {+1: +w, +w: -1, -1: -w, -w: +1}  # +

class Cart:
    def __init__(self, pos, move):
        self.pos = pos
        self.move = move
        self.turn = itertools.cycle([left, straight, right])

carts = []
bypos = {}
for pos, sym in enumerate(tracks):
    if sym in cart_moves:
        cart = Cart(pos, cart_moves[sym])
        carts.append(cart)
        bypos[pos] = cart

crash = None
while crash is None:
    carts.sort(key=operator.attrgetter('pos'))
    for cart in carts:
        track = tracks[cart.pos]
        if track == '\\':
            cart.move = lcurve[cart.move]
        elif track == '/':
            cart.move = rcurve[cart.move]
        elif track == '+':
            cart.move = next(cart.turn)[cart.move]
        del bypos[cart.pos]
        cart.pos += cart.move
        if cart.pos in bypos:
            # Collision!
            crash = cart.pos
            break
        bypos[cart.pos] = cart

x = crash % w
y = crash // w
print(str(x) + ',' + str(y))
