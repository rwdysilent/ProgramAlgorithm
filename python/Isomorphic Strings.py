# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-19 19:44

# leetcode 205
class Solution(object):
    def isIsomorphic(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        s_map = {}
        s_list = []

        for idx, value in enumerate(s):
            if value in s_map:
                s_list.append(s_map[value])
            else:
                s_map[value] = idx
                s_list.append(idx)

        t_map = {}
        t_list = []
        for idx, value in enumerate(t):
            if value in t_map:
                t_list.append(t_map[value])
            else:
                t_map[value] = idx
                t_list.append(idx)

        if s_list == t_list:
            return True

        return False