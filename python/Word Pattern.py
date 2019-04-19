# -*- coding: utf-8 -*-

# leetcode 290
class Solution(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        p_map = {}
        p_list = []
        s_map = {}
        s_list = []

        for idx, value in enumerate(pattern):
            if value in p_map:
                p_list.append(p_map[value])
            else:
                p_map[value] = idx
                p_list.append(idx)

        for idx, value in enumerate(str.split()):
            if value in s_map:
                s_list.append(s_map[value])
            else:
                s_map[value] = idx
                s_list.append(idx)

        if p_list == s_list:
            return True

        return False


class Solution2(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """

        # 大神巧用map,学习
        s = pattern
        t = str.split()
        print(map(s.find, s) == map(t.index, t))
