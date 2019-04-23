# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-23 14:40

# LeetCode 454
class Solution(object):
    def fourSumCount(self, A, B, C, D):
        """
        :type A: List[int]
        :type B: List[int]
        :type C: List[int]
        :type D: List[int]
        :rtype: int
        """

        record = {}
        for i in range(len(C)):
            for j in range(len(D)):
                if C[i] + D[j] in record:
                    record[C[i] + D[j]] += 1
                else:
                    record[C[i] + D[j]] = 1

        res = 0
        for i in range(len(A)):
            for j in range(len(B)):
                if 0 - A[i] - B[j] in record:
                    res += record[0 - A[i] - B[j]]

        return res
