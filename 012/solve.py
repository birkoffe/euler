#!usr/bin/python

def D(a, b):
    count = 0

    f1 = []
    f2 = []
    for i in xrange(1, b/2):
        if a%i == 0:
            f1 += [i]
        if b%i == 0:
            f2 += [i]
    f1 += [a]
    f2 += [b]

    for f in f1:
        if f not in f2 or f == 1:
            count += len(f2)

    return count

i = 0
a, b = 1, 1
while True:
    i += 1
    dd = D(a, b)
    if dd > 500:
        print(dd, a*b)
        break
    if i%2 == 0:
        a += 1
    else:
        b += 2
