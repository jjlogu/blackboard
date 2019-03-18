#!/bin/python3
"""
Constraints:
2<=numOfParts<=10^6
1<=parts[i]<=10^6
0<=i<numOfparts
input: 
8,4,6,12
output:
58

Explanation:
4+6 =10
10+8=18
18+12=30

10+18+30 = 58
"""
from collections import deque
def minimumTime(numOfParts, parts):
    if 2 > numOfParts:
        return parts[0]
    parts.sort()
    dparts = deque(parts)
    total_time = combined = 0
    while numOfParts > 1:
        combined = dparts[0] + dparts[1]
        total_time += combined
        flag = True
        for i in range(2, numOfParts):
            if  combined < dparts[i]:
                dparts.insert(i-1, combined)
                flag = False
        if flag:
            dparts.append(combined)

        dparts.popleft()
        dparts.popleft()

        numOfParts -= 1
        # print(dparts)
    return total_time

if "__main__" == __name__:
    print(minimumTime(4, [8,4,6,12]))
