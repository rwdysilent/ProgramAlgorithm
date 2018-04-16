# encoding: utf-8
# @author: pfwu
# @Date: 2018-04-10 23:19

# https://leetcode.com/problems/linked-list-cycle/description/
#
# Given a linked list, determine if it has a cycle in it.
#
# Follow up:
# Can you solve it without using extra space
# 思路：设置两个指针，一个快指针，一个慢指针，快指针每次移动2个位置，慢指针每次移动1个位置，如果是回环链表，俩指针必定会相遇

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def hasCycle(self, head):
        if head is None:
            return False

        # function 1
        slow = head
        fast = head.next
        while fast is not slow:
            if fast is None or fast.next is None:
                return False
            fast = fast.next.next
            slow = slow.next

        return True

        # function 2
        # slow = fast = head
        # while fast and fast.next:
        #     slow = slow.next
        #     fast = fast.next.next
        #     if fast is slow:
        #         return True
        #
        # return False

