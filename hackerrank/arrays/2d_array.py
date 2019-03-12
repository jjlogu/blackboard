#!/bin/python3
# problem statement: https://www.hackerrank.com/rest/contests/master/challenges/2d-array/download_pdf?language=English

import math
import os
import random
import re
import sys

# Complete the hourglassSum function below.
def hourglassSum(arr):
    max_hour_glass_sum = -63
    for i in range(4):
        for j in range(4):
            top = sum(arr[i][j:j+3])
            mid = arr[i+1][j+1]
            bottom = sum(arr[i+2][j:j+3])
            hour_glass_sum = top+mid+bottom
            if max_hour_glass_sum < hour_glass_sum:
                max_hour_glass_sum = hour_glass_sum
    return max_hour_glass_sum

if __name__ == '__main__':
    arr = []

    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))

    result = hourglassSum(arr)
    print(result)

