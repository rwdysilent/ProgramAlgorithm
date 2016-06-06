# -*- coding: utf-8 -*-
class Solution(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        #myself#
        #d = []
        #p = list(pattern)
        #s = str.split(' ')
        #p.sort()
        #s.sort()
        #i = 0
        #lp = len(p)
        #ls = len(s)
        #if lp != ls:
        #    print False
        #else:
        #    while i < lp - 1:
        #        if p[i] == p[i+1] and s[i] == s[i+1]:
        #            d.append(True)
        #        elif p[i] != p[i+1] and s[i] != s[i+1]:
        #            d.append(True)
        #        else:
        #            d.append(False)
        #        i += 1
        #    if False in d:
        #        print False
        #    else:
        #        print True

        #大神巧用map,学习
        s = pattern
        t = str.split()
        print map(s.find, s) == map(t.index, t)

P = Solution()
P.wordPattern('abba', "dog cat cat dog")
P.wordPattern(pattern = "abba", str = "dog cat cat fish")
P.wordPattern(pattern = "aaaa", str = "dog cat cat dog")
P.wordPattern(pattern = "abba", str = "dog dog dog dog")
P.wordPattern(pattern = "aaa", str = "aa aa aa aa")



#pattern = "abba", str = "dog cat cat dog" should return true.
#pattern = "abba", str = "dog cat cat fish" should return false.
#pattern = "aaaa", str = "dog cat cat dog" should return false.
#pattern = "abba", str = "dog dog dog dog" should return false.