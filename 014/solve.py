#!/usr/bin/python

longestN = 0
longest  = 0
limit = 1000001
hist = [0 for x in range(limit)]

for x in range(2, limit):
    n = x
    i = 0
    while n != 1:
        if n < limit and hist[n] > 0:
            i += hist[n]
            break
        if n%2 == 0:
            n //= 2
        else:
            n = 3*n+1
        i += 1
    hist[x] = i
    if i > longest:
        longest = i
        longestN = x

print(longestN)
