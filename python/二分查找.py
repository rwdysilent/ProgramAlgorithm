# coding: utf-8
# Copyright 2019 pfwu. All rights reserved.
# @Author: pfwu
# @Date: 2019-04-12 15:09


def binary_search(arr, n, target):
    left, right = 0, n - 1
    while left <= right:
        mid = left + (right - left) / 2
        if arr[mid] == target:
            return mid

        if arr[mid] > target:
            right = mid - 1
        else:
            left = mid + 1

    return -1
