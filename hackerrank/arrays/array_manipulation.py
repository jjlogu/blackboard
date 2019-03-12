#!/bin/python3
# problem: https://www.hackerrank.com/rest/contests/master/challenges/crush/download_pdf?language=English

import math
import os
import random
import re
import sys

# Complete the arrayManipulation function below.
def arrayManipulation(n, queries):
    arr = [0] * n
    for query in queries:
        for i in  range(query[0]-1, query[1]):
            arr[i] += query[2]
    return max(arr)

if __name__ == '__main__':
    # fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nm = input().split()

    n = int(nm[0])

    m = int(nm[1])

    queries = []

    for _ in range(m):
        queries.append(list(map(int, input().rstrip().split())))

    result = arrayManipulation(n, queries)
    print(result)

    # fptr.write(str(result) + '\n')

    # fptr.close()

