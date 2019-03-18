#!/bin/python3
# https://www.hackerrank.com/rest/contests/master/challenges/frequency-queries/download_pdf?language=English

import math
import os
import random
import re
import sys

# Complete the freqQuery function below.
def freqQuery(queries):
    ds = {}
    result = []
    for query in queries:
        typ = query[0]
        val = query[1]
        if 1 == typ:
            ds[val] = ds.get(val, 0) + 1
        elif 2 == typ:
            ds[val] = ds.get(val, 0) - 1
            if 0 >= ds[val]:
                del(ds[val])
        else:
            if val in ds.values():
                result.append(1)
            else:
                result.append(0)
    return result

if __name__ == '__main__':
    # fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input().strip())

    queries = []

    for _ in range(q):
        queries.append(list(map(int, input().rstrip().split())))

    ans = freqQuery(queries)
    print(ans)

    # fptr.write('\n'.join(map(str, ans)))
    # fptr.write('\n')

    # fptr.close()
