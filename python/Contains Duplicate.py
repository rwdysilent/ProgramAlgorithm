# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-23 18:42

# LeetCode 217
class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """

        if len(nums) < 2:
            return False

        record = {}
        for i in range(len(nums)):
            if nums[i] in record:
                return True

            record[nums[i]] = i

        return False