# LeetCode 350
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
        return ans


A = [1, 2, 2, 1]
B = [2, 2]
P = Solution()
P.intersect(A, B)


class Solution2(object):
    def intersect(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        d1 = {}
        res = []
        for n in nums1:
            if n not in d1:
                d1[n] = 1
            else:
                d1[n] += 1

        print(d1)
        for n in nums2:
            if n in d1 and d1[n] > 0:
                res.append(n)
                d1[n] -= 1

        return res

A = [1, 2, 2, 1]
B = [2, 2, 3]
P = Solution2()
print(P.intersect(A, B))