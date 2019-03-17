#!/bin/python3

"""
Merge k sorted linked list and return 1 sorted list.


input:
    [
    1->4->5,
    1->3->4,
    2->6
    ]
output:
    1->1->2->3->4->4->5->6
"""

import copy
# from heapq import heapify, heappush, heappop


class Node(object):
    def __init__(self, value, next=None):
        self.value=value
        self.next=next

    def __lt__(self, other):
        return self.value < other.value

    def __repr__(self):
        return "Node({})".format(self.value)

def print_ll(head):
    while head:
        print(head.value, end = "")
        head = head.next

        if head:
            print(" -> ", end="")
        else:
            print("")

def merge_k_sorted_ll(lists):
    head = tail = None
    count = 0
    count1 = 0
    while lists:
        count += 1
        for i in range(1, len(lists)):
            count1 +=1
            if lists[0] > lists[i]:
                lists[0], lists[i] = lists[i], lists[0]
        if not head:
            head = tail = lists[0]
            if lists[0].next:
                lists[0] = lists[0].next
            else:
                del(lists[0])
            tail.next = None

        else:
            tail.next = lists[0]
            if lists[0].next:
                lists[0] = lists[0].next
            else:
                del(lists[0])
            tail = tail.next
    print(count, count1)
    return head 

ll1 = Node(1, Node(4, Node(5)))
ll2 = Node(1, Node(3, Node(4)))
ll3 = Node(2, Node(6))

merged_ll = merge_k_sorted_ll([ll1, ll2, ll3])
merged_ll1 = merge_k_sorted_ll([])

print_ll(merged_ll)
print_ll(merged_ll1)
