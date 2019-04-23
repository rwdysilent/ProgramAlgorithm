# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-23 15:39

# leetcode 447
class Solution(object):
    def numberOfBoomerangs(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """

        res = 0
        for i in points:
            record = {}
            for j in points:
                # if i != j:
                s = i[0] - j[0]
                t = i[1] - j[1]
                key = s * s + t * t
                record[key] = 1 + record.get(key, 0)

            for k in record:
                res += record[k] * (record[k] - 1)

        return res
