# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-15 17:17

# leetcode: 75. Sort Colors
# 三路快排: https://coding.imooc.com/lesson/82.html#mid=2660
# 时间复杂度： O(n)
# 空间复杂度： O(1)


class Solution(object):
    def sortColors(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        zero = -1  # nums[0...zero] == 0
        two = len(nums)  # nums[n-1...two] == 2

        i = 0
        while i < two:
            if nums[i] == 1:
                i += 1
            elif nums[i] == 2:
                two -= 1
                nums[i], nums[two] = nums[two], nums[i]
            else:  # nums[i] == 0
                assert nums[i] == 0
                zero += 1
                nums[i], nums[zero] = nums[zero], nums[i]
                i += 1

# leetcode: 88
# leetcode: 215
