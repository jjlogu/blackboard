#!/bin/python3 

if "__main__" == __name__:
    n = 3
    n = n-1
    inpt = [3, 2,5,6,10,21,15]
    print(inpt)
    for i in range(n):
        if inpt[i] > inpt[n]:
            inpt[i],inpt[n]=inpt[n], inpt[i]
    for i in range(3, len(inpt)):
        if inpt[i] < inpt[n]:
            inpt[i],inpt[n]=inpt[n], inpt[i]
    print(inpt[n])
    print(inpt)
