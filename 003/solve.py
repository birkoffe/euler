#!usr/bin/python

from math import sqrt

def isprime(n):
    for x in xrange(2, int(sqrt(n))+1):
        if n%x == 0:
            return False
    return True

def primes():
    y = 2
    while True:
        pr = y
        yield pr
        y += 1
        while isprime(y) is False:
            y += 1

number = 600851475143
P = primes()

while number != 1:
    x = P.next()
    if number%x == 0:
        number /= x
        result = x

print result
