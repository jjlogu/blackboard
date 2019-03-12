#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the sockMerchant function below.
def sockMerchant(n, ar):
    res = {}
    for color in ar:
        if not res.get(color):
            res[color]=1
        else:
            res[color]+=1
    valid_pairs = 0
    for i in res.values():
        valid_pairs += int(i/2)
    return valid_pairs

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    ar = list(map(int, input().rstrip().split()))

    result = sockMerchant(n, ar)

    fptr.write(str(result) + '\n')

    fptr.close()
