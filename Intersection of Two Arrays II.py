import collections
class Solution(object):
    def intersect(self, nums1, nums2):
        # type: (object, object) -> object
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        c = collections.Counter(nums1)
        ans = []
        for x in nums2:
            if c[x] > 0:
                ans += x,
                c[x] -= 1
        print ans
        return ans


A = [1,2,2,1]
B = [2,2]
P = Solution()
P.intersect(A, B)
