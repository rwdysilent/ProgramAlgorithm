# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-23 18:56

# LeetCode 220
# TODO
class Solution(object):
    def containsNearbyAlmostDuplicate(self, nums, k, t):
        """
        :type nums: List[int]
        :type k: int
        :type t: int
        :rtype: bool
        """

        record = {}
        for i in range(len(nums)):
            # if record[nums[i]] in record:
            #     return True

            record[nums[i]] = i

            if len(record) == k + 1:
                record.pop(nums[i - k])

        return False