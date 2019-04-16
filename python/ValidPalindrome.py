# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-16 16:44

# LeetCode 125
# https://coding.imooc.com/lesson/82.html#mid=2661
class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        i = 0
        j = len(s) - 1

        while i < j:
            if not str.isalnum(str(s[i])):
                i += 1
                continue
            if not str.isalnum(str(s[j])):
                j -= 1
                continue
            if str.lower(str(s[i])) == str.lower(str(s[j])):
                i += 1
                j -= 1
            else:
                return False

        return True


s = "A man, a plan, a canal: Panama"
P = Solution()
print(P.isPalindrome(s))