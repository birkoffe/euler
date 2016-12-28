#!usr/bin/python

max = 1000

for a in xrange(1, max/3):
    for b in xrange(a+1, (max-a)/2):
        c = max-b-a
        if a**2+b**2 == c**2:
            print a*b*c

