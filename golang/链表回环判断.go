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

package main

type ListNodeCycle struct {
	value int
	next  *ListNodeCycle
}

func Solution(head *ListNodeCycle) bool {
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
