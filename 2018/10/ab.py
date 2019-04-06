#!/usr/bin/env python3

import operator
import sys

points = [tuple(map(int, (line[10:16], line[18:24], line[36:38], line[40:42])))
          for line in sys.stdin]

xs, ys, dxs, dys = zip(*points)

# For this puzzle we need a human to look at the sky and determine the answer.
# But we can be smart about *when* the human looks at the sky.  Two assumptions
# make it easier to decide when to look at the sky:
# (1) We assume that the message will appear before any point moves outside of
#     the initial bounds.
# (2) We assume that when the message appears, either the dimensions of the
#     points or the number of distinct values for each dimension of points will
#     be minimized.

class Sky:
    def __init__(self, xs, ys, sec):
        self.xs = xs
        self.ys = ys
        self.sec = sec
        self.min_x = min(xs)
        self.min_y = min(ys)
        self.max_x = max(xs)
        self.max_y = max(ys)
        self.dim_x = self.max_x - self.min_x + 1
        self.dim_y = self.max_y - self.min_y + 1
        self.dis_x = len(set(xs))
        self.dis_y = len(set(ys))
    def next(self):
        next_xs = list(map(operator.add, self.xs, dxs))
        next_ys = list(map(operator.add, self.ys, dys))
        next_sec = self.sec + 1
        return Sky(next_xs, next_ys, next_sec)
    def __str__(self):
        sky = bytearray(ord(c)
                        for _ in range(self.dim_y)
                        for c in '.' * self.dim_x + '\n')
        for x, y in zip(self.xs, self.ys):
            i = (self.dim_x + 1) * (y - self.min_y) + (x - self.min_x)
            sky[i] = ord('#')
        return "Sky after " + str(self.sec) + " seconds:\n" + sky.decode()

sky = Sky(xs, ys, 0)
min_x = sky.min_x
min_y = sky.min_y
max_x = sky.max_x
max_y = sky.max_y
min_dim_x_sky = min_dim_y_sky = min_dis_x_sky = min_dis_y_sky = sky

while (sky.min_x >= min_x and sky.min_y >= min_y and
       sky.max_x <= max_x and sky.max_y <= max_y):
    if sky.dim_x <= min_dim_x_sky.dim_x:
        min_dim_x_sky = sky
    if sky.dim_y <= min_dim_y_sky.dim_y:
        min_dim_y_sky = sky
    if sky.dis_x <= min_dis_x_sky.dis_x:
        min_dis_x_sky = sky
    if sky.dis_y <= min_dis_y_sky.dis_y:
        min_dis_y_sky = sky
    sky = sky.next()

skies = set([min_dim_x_sky, min_dim_y_sky, min_dis_x_sky, min_dis_y_sky])
for sky in sorted(skies, key=operator.attrgetter('sec')):
    print(sky)
