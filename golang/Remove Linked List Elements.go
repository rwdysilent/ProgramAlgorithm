// Copyright 2019 wupengfei. All rights reserved.
//
// @Author: wupengfei
// @Date: 2019-06-10 20:42
//leetcode 203

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	var dummyHead = new(ListNode)
	dummyHead.Next = head
	cur := dummyHead

	if head == nil {
		return nil
	}

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
