# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-17 19:40

# leetcode 202
class Solution(object):
    def isHappy(self, n):
        """
        :type n: int
        :rtype: bool
        """
        assert n >= 0

        start = n
        res = {}
        while start != 1:
            nums = map(lambda x: int(x) * int(x), list(str(start)))
            start = 0
            for i in nums:
                start += int(i)

            if start == 1:
                return True

            if start not in res:
                res[start] = 1
            else:
                return False

        return True


P = Solution()
print(P.isHappy(19))
