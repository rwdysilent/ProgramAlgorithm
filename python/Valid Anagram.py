# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-17 17:55

# LeetCode 242
class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        s_map = {}
        t_map = {}
        for i in s:
            if i not in s_map:
                s_map[i] = 1
            else:
                s_map[i] += 1

        for j in t:
            if j not in t_map:
                t_map[j] = 1
            else:
                t_map[j] += 1

        if s_map == t_map:
            return True

        return False