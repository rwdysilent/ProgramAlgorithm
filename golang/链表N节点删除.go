// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-03-27 11:22

//https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

//Given a linked list, remove the nth node from the end of list and return its head.
//
//For example,
//
//   Given linked list: 1->2->3->4->5, and n = 2.
//
//   After removing the second node from the end, the linked list becomes 1->2->3->5.
//Note:
//Given n will always be valid.
//Try to do this in one pass.

package main

type ListNodeRemove struct {
	Val  int
	Next *ListNodeRemove
}

func removeNthFromEnd(head *ListNodeRemove, n int) *ListNodeRemove {
	var dummy  = new(ListNodeRemove)
	dummy.Next = head
	first := dummy
	second := dummy

	for i := 1; i <= n + 1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}
