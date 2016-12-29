#!/usr/bin/python

def get(i, j, grid, move):
    pr = 1
    if j+move[7] > 19 or i+move[3] > 19:
        return 0
    for x in xrange(4):
        pr *= grid[i+move[x]][j+move[4+x]]
    return pr

grid = []
for l in open("input.txt", "r"):
    grid += [[int(x) for x in l[:-1].split()]]

result = 0

moves = [[0, 0, 0, 0, 0, 1, 2, 3], [0, 1, 2, 3, 0, 0, 0, 0], [0, 1, 2, 3, 0, 1, 2, 3], [0, -1, -2, -3, 0, 1, 2, 3]]

for a in xrange(20):
    for b in xrange(20):
        for move in moves:
            z = get(a, b, grid, move)
            result = max(result, z)

print result
