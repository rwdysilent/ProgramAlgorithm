// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-11 18:42

package main

func binary_search(arr []int, low, high, hkey int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2

	if arr[mid] > hkey {
		return binary_search(arr, low, mid-1, hkey)
	} else if arr[mid] < hkey {
		return binary_search(arr, mid+1, high, hkey)
	}

	return mid
}
