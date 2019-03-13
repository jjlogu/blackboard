#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the twoStrings function below.
def twoStrings1(s1, s2):
    arr ={}
    for i in range(len(s1)):
        arr[s1[i]] = 1
    for i in range(len(s2)):
        if arr.get(s2[i]):
            return("YES")
    return("NO")

# Complete the twoStrings function below.
def twoStrings(s1, s2):
    # create sets of unique characters
    # and test for intersection
    if set(s1)&set(s2):
        return "YES"
    else:
        return "NO"

if __name__ == '__main__':
    # fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())
    import time
    start = time.time()
    for q_itr in range(q):
        s1 = input()

        s2 = input()

        result = twoStrings(s1, s2)

        # fptr.write(result + '\n')

    # fptr.close()
    end = time.time()
    print(end - start)
