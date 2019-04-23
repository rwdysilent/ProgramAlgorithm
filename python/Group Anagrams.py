# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-23 14:57

# leetcode 49
class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """

        record = {}
        for s in strs:
            l = sorted([ord(d) for d in s])
            if str(l) in record:
                record[str(l)].append(s)
            else:
                record[str(l)] = [s]

        return record.values()


ss = ["cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"]
P = Solution()
print(P.groupAnagrams(ss))
