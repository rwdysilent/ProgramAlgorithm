// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-03-18 00:10

package main

import (
	"fmt"
)

//堆排序
func HeapSort(array []int) {
	m := len(array)
	s := m / 2
	for i := s; i > -1; i-- {
		heap(array, i, m-1)
	}
	for i := m - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		heap(array, 0, i-1)
	}
}

func heap(array []int, i, end int) {
	//左子树
	l := 2*i + 1
	if l > end {
		return
	}

	n := l

	//右子数
	r := 2*i + 2
	if r <= end && array[r] > array[l] {
		n = r
	}

	//根节点大于左子树，返回
	if array[i] > array[n] {
		return
	}

	array[n], array[i] = array[i], array[n]

	fmt.Println("arry: ", array)
	heap(array, n, end)
}

func main() {
	dData := []int{
		55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,
	}
	HeapSort(dData)
	fmt.Println(dData)
}
