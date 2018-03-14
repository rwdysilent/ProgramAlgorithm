# refer to https://leetcode.com/discuss/105919/a-clean-python-solution
class Solution(object):
    def addDigits(self, num):
        """
        :type num: int
        :rtype: int
        """
        if num == 0:
            print num
        else:
            tmp = num % 9
            print 9 if tmp == 0 else tmp

P = Solution()
P.addDigits(18)