// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-03-27 11:13

//https://leetcode.com/problems/linked-list-cycle/description/
//
//Given a linked list, determine if it has a cycle in it.
//
//Follow up:
//Can you solve it without using extra space?

//思路：设置两个指针，一个快指针，一个慢指针，快指针每次移动2个位置，慢指针每次移动1个位置，如果是回环链表，俩指针必定会相遇

//此题go语言无法实现，暂时不清楚为什么

package main

import "fmt"

type ListNodeCycle struct {
	value int
	next  *ListNodeCycle
}

func hasCycle(head *ListNodeCycle) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow := head
	fast := head.next

	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}

	return true
}

func main() {
	node1 := new(ListNodeCycle)
	node2 := new(ListNodeCycle)
	node3 := new(ListNodeCycle)
	node4 := new(ListNodeCycle)
	node5 := new(ListNodeCycle)
	node1.value = 1
	node2.value = 2
	node3.value = 3
	node4.value = 4
	node5.value = 5
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node1

	fmt.Println(hasCycle(node1))
}