# coding: utf-8
# Copyright 2019 wupengfei. All rights reserved.
# @Author: wupengfei
# @Date: 2019-05-27 20:27
# LeetCode92

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
        if m == n:
            return head

        dummy = ListNode(0)
        dummy.next = head
        pre = dummy

        i = 0
        while i < m - 1:
            pre = pre.next
            i += 1

        reverse = None
        cur = pre.next
        for i in range(n - m + 1):
            nex = cur.next

            cur.next = reverse
            reverse = cur
            cur = nex

        pre.next.next = cur
        pre.next = reverse

        return dummy.next
