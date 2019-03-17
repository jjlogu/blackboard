#!/bin/python3
"""
You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

Example:

Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7

7243+564 = 7807

"""

class Node(object):
    def __init__(self, val, next = None):
        self.val = val
        self.next = next
    def __repr__(self):
        return "Node({})".format(self.val)

def print_ll(head):
    while head:
        print(head.val, end="")
        if head.next:
            print(" -> ", end = "")
        else:
            print()
        head = head.next

# size(ll1) == size(ll2)
def add(ll1, ll2):
    if ll1 == None:
        return None,0
    result = Node(0)
    result.next, carry = add(ll1.next, ll2.next)
    result.val = ll1.val+ll2.val+carry
    carry=int(result.val/10)
    result.val %= 10

    return result, carry

    

ll1=Node(7,Node(2,Node(4,Node(3))))
ll2=Node(5,Node(6,Node(4,Node(2))))

print_ll(ll1)
print_ll(ll2)
result, carry = add(ll1,ll2)
if carry:
    result = Node(carry, result)
print_ll(result)

