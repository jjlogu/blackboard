#!/bin/python3
# problem statement: https://www.hackerrank.com/rest/contests/master/challenges/ctci-array-left-rotation/download_pdf?language=English

import math
import os
import random
import re
import sys

# Complete the rotLeft function below.
def rotLeft(a, d):
    """
    if len(a) == d:
        return a
    for _ in range(d):
        b = a[1:]
        b.append(a[0])
        a = b
    return a
    """
    i = d % len(a)
    return(a[i:]+a[:i])


if __name__ == '__main__':
    #fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nd = input().split()

    n = int(nd[0])

    d = int(nd[1])

    a = list(map(int, input().rstrip().split()))

    result = rotLeft(a, d)
    print(result)

    #fptr.write(' '.join(map(str, result)))
    #fptr.write('\n')

    #fptr.close()

