# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-21 00:39

# leetcode 15
# TODO
class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """

        rsp_list = []
        for i in range(len(nums) - 2):
            sums = 0 - nums[i]
            find_map = {}
            # print("sums: %s, %s" % (nums[i], sums))
            for j in range(i + 1, len(nums)):
                target = sums - nums[j]
                # print("target: %s, %s" % (nums[j], target))
                if target in find_map:
                    ans_list = [nums[i], nums[find_map[target]], nums[j]]
                    # print("ans_list: %s" % ans_list)
                    ans_list = sorted(ans_list)
                    if ans_list not in rsp_list:
                        rsp_list.append(ans_list)
                    # rsp_list.append(ans_list)
                else:
                    find_map[nums[j]] = j
                # print("map: %s" % find_map)

        return rsp_list


n = [[-1, 0, 1, 2, -1, -4],
     [1, 2, -2, -1],
     [0, 0, 0, 0],
     [3, 0, -2, -1, 1, 2],
     ]
# n = [[3, 0, -2, -1, 1, 2]]
P = Solution()
for l in n:
    print(P.threeSum(l))
