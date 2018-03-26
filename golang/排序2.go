// Copyright 2018 . All rights reserved.
//
// @Author: pfwu
// @Date: 2018-03-15 22:11

package main

import "fmt"

var d = []int{33, 5, 76, 4, 34, 32, 6}

//快速排序
//快速排序一听名字就觉得很高端，在实际应用当中快速排序确实也是表现最好的排序算法。快速排序虽然高端，但其实其思想是来自冒泡排序，冒泡排序是通过
//相邻元素的比较和交换把最小的冒泡到最顶端，而快速排序是比较和交换小数和大数，这样一来不仅把小数冒泡到上面同时也把大数沉到下面。
//举个栗子：对5,3,8,6,4这个无序序列进行快速排序，思路是右指针找比基准数小的，左指针找比基准数大的，交换之。
//
//5,3,8,6,4 用5作为比较的基准，最终会把5小的移动到5的左边，比5大的移动到5的右边。
//
//5,3,8,6,4 首先设置i,j两个指针分别指向两端，j指针先扫描（思考一下为什么？）4比5小停止。然后i扫描，8比5大停止。交换i,j位置。
//
//5,3,4,6,8 然后j指针再扫描，这时j扫描4时两指针相遇。停止。然后交换4和基准数。
//
//4,3,5,6,8 一次划分后达到了左边比5小，右边比5大的目的。之后对左右子序列递归排序，最终得到有序序列。
//
//上面留下来了一个问题为什么一定要j指针先动呢？首先这也不是绝对的，这取决于基准数的位置，因为在最后两个指针相遇的时候，要交换基准数到相遇的位置。一般选取第一个数作为基准数，那么就是在左边，所以最后相遇的数要和基准数交换，那么相遇的数一定要比基准数小。所以j指针先移动才能先找到比基准数小的数。
//
//快速排序是不稳定的，其时间平均时间复杂度是O(nlgn)。
func quickSort(num []int, left, right int) []int {
	if num == nil || len(num) == 0 {
		return num
	}
	if left < right {
		temp := num[left]
		i, j := left, right
		for {
			//从右向左
			for num[j] >= temp && i < j {
				j--
			}

			//从左向右
			for num[i] <= temp && i < j {
				i++
			}

			if i >= j {
				break
			}

			num[i], num[j] = num[j], num[i]

		}
		num[left] = num[i]
		num[i] = temp

		quickSort(num, left, i-1)
		quickSort(num, i+1, right)
	}
	return num
}

//维基百科例子
func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
			fmt.Println("if: ", data)
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
			fmt.Println("else: ", data)
		}
	}
	data[head] = mid
	qsort(data[:head])
	qsort(data[head+1:])
}

//练习
func myQsort(data []int) {
	if len(data) <= 1 {
		return
	}

	base := data[0]
	head, tail := 0, len(data)-1

	for i := 1; i <= tail; {
		if data[i] > base {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}

	data[head] = base
	myQsort(data[:head])
	myQsort(data[head+1:])
}

func main() {
	//fmt.Println(quickSort(d, 0, len(d)-1))
	//fmt.Println(d)
	//fmt.Println("========")
	//qsort(d)
	//fmt.Println("========")
	//fmt.Println(d)
	myQsort(d)
	fmt.Println(d)
}
