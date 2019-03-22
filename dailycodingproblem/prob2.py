"""
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
"""

def product_of_all(numbs):
    p = 1
    zeros = 0
    size = len(numbs)
    for i in numbs:
        if i==0:
            zeros += 1
            if zeros == 2:
                return [0] * size
            else:
                continue
        p *= i
    r = []
    if zeros == 0:
        for i in numbs:
            r.append(int(p/i))
    else:
        for i in numbs:
            if 0 == i:
                r.append(p)
            else:
                r.append(0)
    return r

if "__main__" == __name__:
    print(product_of_all([0,0,3,4,5]))
    print(product_of_all([0,2,3,4,5]))
    print(product_of_all([1,2,3,4,5]))
    print(product_of_all([-1,2,-3,4,5]))
