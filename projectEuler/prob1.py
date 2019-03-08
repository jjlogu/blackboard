#!/usr/bin/env python3
'''
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
'''

result = set()

for i in range(1,334):
    result.add(i*3)

for i in range(1,200):
    result.add(i*5)

sum = 0

for i in result:
    sum = sum + i

print(sum)
