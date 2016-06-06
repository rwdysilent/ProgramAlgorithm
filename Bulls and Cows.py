# -*- coding: utf-8 -*-
from collections import Counter
import operator
class Solution(object):
    def getHint(self, secret, guess):
        """
        :type secret: str
        :type guess: str
        :rtype: str
        """
        #u'参考大神的方法'
        #https://leetcode.com/discuss/67037/python-3-lines-solution
        s, g = Counter(secret), Counter(guess)
        a = sum(i == j for i, j in zip(secret, guess))
        print '%sA%sB' % (a, sum((s & g).values()) - a)

        #u'另一位大神的方法'
        #https://leetcode.com/discuss/67016/3-lines-in-python
        #bulls = sum(map(operator.eq, secret, guess))
        #both = sum(min(secret.count(x), guess.count(x)) for x in '0123456789')
        #print '%dA%dB' % (bulls, both - bulls)

P = Solution()
P.getHint(secret="1123", guess="0111")
P.getHint(secret="1807", guess="7810")
P.getHint(secret="1122", guess="2211")
