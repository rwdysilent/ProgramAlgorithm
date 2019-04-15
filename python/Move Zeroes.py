# leetcode 283. Move Zeroes
class Solution1(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        nums.sort(key=lambda x: x == 0)
        return nums


P = Solution1()
print("Solution1: ", P.moveZeroes(nums=[0, 1, 0, 3, 12]))


# given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].


class Solution2(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        k = 0
        for i in range(len(nums)):
            if nums[i]:
                if k != i:
                    nums[k], nums[i] = nums[i], nums[k]
                k += 1

        return nums


P = Solution2()
print("Solution2: ", P.moveZeroes(nums=[0, 1, 0, 3, 12]))
