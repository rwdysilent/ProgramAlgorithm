# coding: utf-8
# Copyright 2019 wupengfei. All rights reserved.
# @Author: wupengfei
# @Date: 2019-05-25 11:37
# leetcode 206

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """

        pre = None
        cur = head

        while cur is not None:
            next = cur.next

            cur.next = pre
            pre = cur
            cur = next

        return pre