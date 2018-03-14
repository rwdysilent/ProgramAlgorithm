# -*- coding: utf-8 -*-
#极简方法
import re
class Solution(object):
    def reverseVowels(self, s):
        """
        :type s: str
        :rtype: str
        """
        #print re.sub('(?i)[aeiou]', lambda m, v=re.findall('(?i)[aeiou]', s): v.pop(), s)

        #vowels = re.findall('(?i)[aeiou]', s)
        #print re.sub('(?i)[aeiou]', lambda m: vowels.pop(), s)
        vowels = re.findall('[aeiou]', s, flags = re.IGNORECASE)
        print re.sub('[aeiou]', lambda m: vowels.pop(), s, flags = re.IGNORECASE)

#传统方法
class Solution2(object):
    def reverseVowels(self, s):
        vowels = "aeiouAEIOU"
        s = list(s)
        ptr_1, ptr_2 = 0, len(s) - 1
        while ptr_1 < ptr_2:
            if s[ptr_1] in vowels and s[ptr_2] in vowels:
                s[ptr_1], s[ptr_2] = s[ptr_2], s[ptr_1]
                ptr_1 += 1
                ptr_2 -= 1
            if s[ptr_1] not in vowels:
                ptr_1 += 1
            if s[ptr_2] not in vowels:
                ptr_2 -= 1
        print ''.join(s)

P = Solution2()
P.reverseVowels("hello")
