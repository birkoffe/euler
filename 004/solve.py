#!usr/bin/python

def palindrome(x):
    result = 0
    while x > 0:
        result = result*10+x%10
        x /= 10
    return result

result = 0
for i in xrange(1, 1000):
    for j in xrange(i, 1000):
        ij = i*j
        if ij == palindrome(ij) and ij > result:
            result = ij

print result