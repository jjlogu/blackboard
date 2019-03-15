#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the countTriplets function below.
def countTriplets(arr, r):
    left = {}
    right = {}
    answer = 0
    for i in arr:
       right[i] = right.get(i, 0) + 1
    for i in arr:
        right[i] -= 1
        if 0 == i % r:
           answer += left.get(i/r, 0) * right.get(i*r, 0)
        left[i] = left.get(i, 0) + 1
    return answer

if __name__ == '__main__':
    # fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nr = input().rstrip().split()

    n = int(nr[0])

    r = int(nr[1])

    arr = list(map(int, input().rstrip().split()))

    ans = countTriplets(arr, r)
    print(ans)

    # fptr.write(str(ans) + '\n')

    # fptr.close()
