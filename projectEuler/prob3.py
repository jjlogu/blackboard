#!/usr/bin/env python3
'''
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
'''
import math

def largest_prime(n):
    lprime = -1
    # For even number largest prime number is 2
    if 0 == n%2:
        lprime = 2
        n = n/2
    # Since we are looking for prime number looping till sqrt(n) is enough
    # At this point factor will be an odd number
    for i in  range(3, math.ceil(math.sqrt(n)), 2):
        if 0 == n%i:
            lprime = i
            n = n/i

    # if the given number itself a prime number
    if n>2:
        lprime = n  
    return lprime

print(largest_prime(600851475143))
