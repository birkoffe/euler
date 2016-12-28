#!usr/bin/python

ssq = 0
sqs = 0

for x in xrange(1,101):
    ssq += x**2
    sqs += x

print sqs**2-ssq
