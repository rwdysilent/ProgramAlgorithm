# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-16 11:26

# leetcode 167
# https://coding.imooc.com/lesson/82.html#mid=2661
class Solution(object):
    def twoSum(self, numbers, target):
        """
        :type numbers: List[int]
        :type target: int
        :rtype: List[int]
        """
        assert len(numbers) >= 2

        for i in range(len(numbers)):
            t = target - numbers[i]
            left = i
            right = len(numbers) - 1
            while left <= right:
                mid = left + (right - left) / 2
                if numbers[mid] == t:
                    if mid == i:
                        return i + 1, i + 2
                    else:
                        return i + 1, mid + 1
                if numbers[mid] < t:
                    left = mid + 1
                else:
                    right = mid - 1
        return -1


nums = [1, 2, 3, 4, 4, 9, 56, 90]
P = Solution()
print(P.twoSum(nums, 8))


#双指针解法
class Solution2(object):
    def twoSum(self, numbers, target):
        """
        :type numbers: List[int]
        :type target: int
        :rtype: List[int]
        """
        assert len(numbers) >= 2
        i = 0
        j = len(numbers) - 1
        while i < j:
            if numbers[i] + numbers[j] == target:
                return i + 1, j + 1
            elif numbers[i] + numbers[j] > target:
                j -= 1
            else:
                i += 1

        return -1


nums = [1, 2, 3, 4, 4, 9, 56, 90]
P = Solution2()
print(P.twoSum(nums, 8))
