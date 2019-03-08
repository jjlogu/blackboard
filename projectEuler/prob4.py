#!/usr/bin/env python3

'''
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
'''

lpalindrome = 0
for i in range(999, 99, -1):
    for j in range(i, 99, -1):
        multiply = str(i*j)
        if multiply == multiply[::-1]:
            if int(multiply) > lpalindrome:
                lpalindrome = int(multiply)
            break

print(lpalindrome)
