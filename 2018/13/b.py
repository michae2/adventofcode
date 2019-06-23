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
        self.crashed = False
    def __str__(self):
        return str(self.pos)

carts = []
bypos = {}
for pos, sym in enumerate(tracks):
    if sym in cart_moves:
        cart = Cart(pos, cart_moves[sym])
        carts.append(cart)
        bypos[pos] = cart

while len(carts) > 1:
    carts.sort(key=operator.attrgetter('pos'))
    for cart in carts:
        if cart.crashed:
            continue
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
            # Collision!  Remove the other cart as well.
            cart.crashed = True
            cart2 = bypos.pop(cart.pos)
            cart2.crashed = True
        else:
            bypos[cart.pos] = cart
    carts = [cart for cart in carts if not cart.crashed]

assert len(carts) == 1
x = carts[0].pos % w
y = carts[0].pos // w
print(str(x) + ',' + str(y))
