#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the checkMagazine function below.
def checkMagazine(magazine, note):
    arr = {}
    for word in magazine:
        if arr.get(word):
            arr[word] += 1
        else:
            arr[word] = 1
    for word in note:
        if arr.get(word):
            arr[word] -= 1
            if 0 > arr[word]:
                print("No")
                return
        else:
           print("No")
           return
    print("Yes")

if __name__ == '__main__':
    mn = input().split()

    m = int(mn[0])

    n = int(mn[1])

    magazine = input().rstrip().split()

    note = input().rstrip().split()

    checkMagazine(magazine, note)
