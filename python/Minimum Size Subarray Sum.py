# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-16 17:59

# LeetCode 209
# https://coding.imooc.com/lesson/82.html#mid=2662
class Solution(object):
    def minSubArrayLen(self, s, nums):
        """
        :type s: int
        :type nums: List[int]
        :rtype: int
        """

        i = 0
        j = -1
        sums = 0
        res = len(nums) + 1

        while i < len(nums):
            if sums < s and j + 1 < len(nums):
                j += 1
                sums += nums[j]
            else:
                sums -= nums[i]
                i += 1

            if sums >= s:
                if res > j - i + 1:
                    res = j - i + 1

        if res == len(nums) + 1:
            return 0
        return res


target = 4
num = [1, 4, 4]
P = Solution()
print(P.minSubArrayLen(target, num))
