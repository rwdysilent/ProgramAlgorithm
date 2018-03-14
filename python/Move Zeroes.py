class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        nums.sort(key=lambda x: x == 0)
        print nums



P = Solution()
P.moveZeroes(nums = [0, 1, 0, 3, 12])
#given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].