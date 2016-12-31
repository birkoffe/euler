#!/usr/bin/python

from math import log10

result = 0
for x in open("input.txt", "r"):
    result += int(x)

print result/(10**(int(log10(result))-9))
