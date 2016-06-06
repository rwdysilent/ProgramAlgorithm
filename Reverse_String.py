class Solution(object):
    def reverseString(self, s):
        """
        :type s: str
        :rtype: str
        """
        L = list(s)
        L.reverse()
        print ''.join(L)
        return ''.join(L)

P = Solution()
P.reverseString('hello')