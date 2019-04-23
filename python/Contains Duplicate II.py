# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-23 17:26

# LeetCode 219
class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """

        record = {}
        for i in range(len(nums)):
            if record.get(nums[i]):
                return True

            record[nums[i]] = True

            if len(record) == k + 1:
                record.pop(nums[i - k])

        return False

n = [1, 2, 3, 1, 2, 3]
n = [1, 2, 3, 1]
n = [1, 0, 1, 1]
n = [2, 2]
n = [1, 2, 3, 1]
P = Solution()
print(P.containsNearbyDuplicate(n, 3))


class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        record = {}
        for i in range(len(nums)):
            if nums[i] in record and i - record[nums[i]] <= k:
                return True
            record[nums[i]] = i

        return False
