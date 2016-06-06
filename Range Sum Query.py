class NumArray(object):
    def __init__(self, nums):
        """
        initialize your data structure here.
        :type nums: List[int]
        """
        self.nums = nums

    def sumRange(self, i, j):
        """
        sum of elements nums[i..j], inclusive.
        :type i: int
        :type j: int
        :rtype: int
        """
        initN = 0
        for n in range(i, j+1):
            initN += self.nums[n]
        print initN



numArray = NumArray([-2, 0, 3, -5, 2, -1])
numArray.sumRange(0,5)
        # Your NumArray object will be instantiated and called as such:
        # numArray = NumArray(nums)
        # numArray.sumRange(0, 1)
        # numArray.sumRange(1, 2)