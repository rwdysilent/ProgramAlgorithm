# coding: utf-8
# Copyright 2019 wupengfei. All rights reserved.
# @Author: wupengfei
# @Date: 2019-05-27 20:27
# LeetCode92

# Todo
# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def reverseBetween(self, head, m, n):
        """
        :type head: ListNode
        :type m: int
        :type n: int
        :rtype: ListNode
        """

        pre = ListNode(0)
        pre.next = head
        i = 0
        while i < m:
            pre = pre.next
            i += 1

        head = pre.next
        for i in range(m, n):
            while i < n:
                nex = head.next

                head.next = head.next.next
                nex.next = pre.next
                pre.next = nex
                i += 1

        return pre
