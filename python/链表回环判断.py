# encoding: utf-8
# @author: pfwu
# @Date: 2018-04-10 23:19


class Solution(object):
    def hasCycle(self, head):
        if head is None:
            return False

        # function 1
        slow = head
        fast = head.next
        while fast is not slow:
            if fast is None or fast.next is None:
                return False
            fast = fast.next.next
            slow = slow.next

        return True

        # function 2
        # slow = fast = head
        # while fast and fast.next:
        #     slow = slow.next
        #     fast = fast.next.next
        #     if fast is slow:
        #         return True
        #
        # return False

