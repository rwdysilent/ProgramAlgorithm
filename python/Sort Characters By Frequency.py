# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-19 19:54

# leetcode 451
class Solution(object):
    def frequencySort(self, s):
        """
        :type s: str
        :rtype: str
        """
        s_map = {}
        for i in s:
            if i not in s_map:
                s_map[i] = 1
            else:
                s_map[i] += 1

        s_values = s_map.values()
        s_values.sort()

        res = ''
        for i in reversed(s_values):
            for k, v in s_map.items():
                if i == v:
                    print(k)
                    for _ in range(v):
                        res += k
                    s_map.pop(k)
        return res


s = "tree"
P = Solution()
print(P.frequencySort(s))
