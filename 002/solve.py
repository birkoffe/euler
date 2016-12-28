#!usr/bin/python

S, x, y = 0, 1, 2
while S < 4000000:
    if x%2 == 0:
        S += x
    x, y = y, x+y

print S
