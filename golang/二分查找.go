// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-11 18:42

// 二分搜索（英语：binary search），也称折半搜索（英语：half-interval search）[1]、对数搜索（英语：logarithmic search）[2]，
// 是一种在有序数组中查找某一特定元素的搜索算法。搜索过程从数组的中间元素开始，如果中间元素正好是要查找的元素，则搜索过程结束；如果某一
// 特定元素大于或者小于中间元素，则在数组大于或小于中间元素的那一半中查找，而且跟开始一样从中间元素开始比较。如果在某一步骤数组为空，则
// 代表找不到。这种搜索算法每一次比较都使搜索范围缩小一半。

package main

import "fmt"

func binarySearch(arr []int, low, high, hkey int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2

	if arr[mid] > hkey {
		return binarySearch(arr, low, mid-1, hkey)
	} else if arr[mid] < hkey {
		return binarySearch(arr, mid+1, high, hkey)
	}

	return mid
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	n := binarySearch(arr, 0, 7, 4)
	fmt.Println(n)
}
