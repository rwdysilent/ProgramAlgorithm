# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-16 18:51

# leetcode 3
# https://coding.imooc.com/lesson/82.html#mid=2663
class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        i = 0
        j = -1
        res = 0
        nums = {}

        while i < len(s):
            if j + 1 < len(s) and s[j + 1] not in nums:
                nums[s[j + 1]] = 0

            if j + 1 < len(s) and nums[s[j + 1]] == 0:
                j += 1
                nums[s[j]] += 1
            else:
                nums[s[i]] -= 1
                i += 1

            res = max(res, (j - i + 1))
        return res
