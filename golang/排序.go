// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-03-14 21:36

package main

import (
	"fmt"
)

var data = []int{33, 5, 76, 4, 34, 32, 6}

//冒泡排序 时间复杂度O(n^2)
//冒泡排序是最简单的排序之一了，其大体思想就是通过与相邻元素的比较和交换来把小的数交换到最前面。这个过程类似于水泡向上升一样，
//因此而得名。举个栗子，对5,3,8,6,4这个无序序列进行冒泡排序。首先从后向前冒泡，4和6比较，把4交换到前面，序列变成5,3,8,4,6。
//同理4和8交换，变成5,3,4,8,6,3和4无需交换。5和3交换，变成3,5,4,8,6,3.这样一次冒泡就完了，把最小的数3排到最前面了。对剩下
//的序列依次冒泡就会得到一个有序序列。冒泡排序的时间复杂度为O(n^2)。
func maoPao(num []int) []int {
	if num == nil || len(num) == 1 {
		return num
	}
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
	fmt.Println(num)
	return num
}

//选择排序
//选择排序的思想其实和冒泡排序有点类似，都是在一次排序后把最小的元素放到最前面。但是过程不同，冒泡排序是通过相邻的比较和交换。
//而选择排序是通过对整体的选择。举个栗子，对5,3,8,6,4这个无序序列进行简单选择排序，首先要选择5以外的最小数来和5交换，也就是
//选择3和5交换，一次排序后就变成了3,5,8,6,4.对剩下的序列一次进行选择和交换，最终就会得到一个有序序列。其实选择排序可以看成
//冒泡排序的优化，因为其目的相同，只是选择排序只有在确定了最小数的前提下才进行交换，大大减少了交换的次数。选择排序的时间复杂度
//为O(n^2)
func selectSort(num []int) []int {
	if num == nil || len(num) == 1 {
		return num
	}
	for i := 0; i < len(num); i++ {
		minIndex := i
		for j := i + 1; j < len(num); j++ {
			if num[j] < num[minIndex] {
				minIndex = j
			}
		}
		num[i], num[minIndex] = num[minIndex], num[i]
	}
	fmt.Println(num)
	return num
}

//插入排序
//插入排序不是通过交换位置而是通过比较找到合适的位置插入元素来达到排序的目的的。相信大家都有过打扑克牌的经历，特别是牌数较大
//的。在分牌时可能要整理自己的牌，牌多的时候怎么整理呢？就是拿到一张牌，找到一个合适的位置插入。这个原理其实和插入排序是一样
//的。举个栗子，对5,3,8,6,4这个无序序列进行简单插入排序，首先假设第一个数的位置时正确的，想一下在拿到第一张牌的时候，没必要
//整理。然后3要插到5前面，把5后移一位，变成3,5,8,6,4.想一下整理牌的时候应该也是这样吧。然后8不用动，6插在8前面，8后移一位，
//4插在5前面，从5开始都向后移一位。注意在插入一个数的时候要保证这个数前面的数已经有序。简单插入排序的时间复杂度也是O(n^2)。
func insertSort(num []int) []int {
	if num == nil || len(num) == 1 {
		return num
	}
	for i := 1; i < len(num); i++ {
		for j := i; j > 0; j-- {
			if num[j-1] > num[j] {
				num[j], num[j-1] = num[j-1], num[j]
			}
		}
		fmt.Println(num)
	}
	fmt.Println(num)
	return num
}

func main() {
	//maoPao(data)
	//selectSort(data)
	insertSort(data)
}
