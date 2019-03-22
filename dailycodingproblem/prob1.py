"""
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
"""

def add_up_to_k(numbs, k):
    numbs.sort()
    size = len(numbs)
    for i in range(size):
        if numbs[i] > k:
            return False
        for j in range(i+1, size):
            if numbs[i] + numbs[j] == k:
                return True
    return False

if __name__ == "__main__":
    print(add_up_to_k([10,15,3,7], 19))
    print(add_up_to_k([], 19))
