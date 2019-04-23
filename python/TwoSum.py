# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-19 23:57

# LeetCode 1
# 构建查找表
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        assert len(nums) >= 2
        map_nums = {}
        # for i, d in enumerate(nums):
        #     if target - d in map_nums:
        #         return [map_nums[target - d], i]
        #     else:
        #         map_nums[d] = i

        for i in range(len(nums)):
            if target - nums[i] in map_nums:
                return [map_nums[target - nums[i]], i]
            else:
                map_nums[nums[i]] = i


# numbers = [3, 3]
numbers = [1, 2, -2, -1]
P = Solution()
print(P.twoSum(numbers, 1))
