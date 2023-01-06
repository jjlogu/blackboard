"""
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
"""
# https://leetcode.com/problems/add-two-numbers/
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        result_head=None
        result=result_head
        carry_forwader = 0
        while l1:
            sum_val = l1.val+carry_forwader
            if l2:
                sum_val += l2.val
                l2=l2.next
            carry_forwader=int(sum_val/10)
            if result:
                result.next=ListNode(sum_val%10)
                result=result.next
            else:
                result_head=ListNode(sum_val%10)
                result=result_head
            l1=l1.next
        while l2:
            sum_val = l2.val+carry_forwader
            result.next=ListNode(sum_val%10)
            result=result.next
            l2=l2.next
        if carry_forwader != 0:
            result.next=ListNode(carry_forwader)
        return result_head



