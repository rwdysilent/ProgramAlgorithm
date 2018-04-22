// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-11 18:42

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
	arr := []int{1, 3, 2, 5, 4, 7, 6}

	n := binarySearch(arr, 0, 7, 7)
	fmt.Println(n)
}
