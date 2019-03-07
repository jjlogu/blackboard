#!/usr/bin/env python3

#    Pythagorean Triplet in an array
#    Given an array of integers, write a function that returns true if there is a triplet (a, b, c) that satisfies a2 + b2 = c2.
#    
#    Example:
#
#    Input: arr[] = {3, 1, 4, 6, 5}
#    Output: True
#    There is a Pythagorean triplet (3, 4, 5).
#    
#    Input: arr[] = {10, 4, 6, 12, 5}
#    Output: False
#    There is no Pythagorean triplet.

def has_pythagorean_triplet(input):
    square = list(map(lambda x: x*x, input))
    square.sort()
    n = len(square)
    '''
    for i in range(n-1, 1, -1):
        for j in range(0, i-1):
            for k in range(j+1, i):
                print(i,j,k)
                if square[i] == square[j] + square[k]:
                    print(square[i], square[j], square[k])
                    return True
    '''
    for i in range(n-1, 1, -1):
        j = 0
        k = i-1
        while( j < k):
            print(i,j,k)
            if square[i] == square[j]+square[k]:
                print(square[i], square[j], square[k])
                return True
            else:
                if square[i] > square[j]+square[k]: 
                    j = j + 1
                else:
                    k = k - 1
    return False
if __name__ == "__main__":
    input = [3, 1, 4, 6, 5, 7, 9]
    print(has_pythagorean_triplet(input))

