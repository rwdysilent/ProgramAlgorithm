// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-03-24 22:00
//leetcode: https://leetcode.com/problems/reverse-linked-list/description/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}

    var pre *ListNode
    cur := head

    for cur != nil {
        next := cur.Next

        cur.Next = pre
        pre = cur
        cur = next
    }

    return pre
}

/**
 * Definition for singly-linked list.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	saveNext := head.Next
	var next *ListNode = nil
	for saveNext != nil {
		head.Next = next
		next = head
		head = saveNext
		saveNext = head.Next
		if saveNext == nil {
			head.Next = next
			return head
		}
	}

	return head
}

//练习
func myReverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}

	q := head.Next
	var next *ListNode
	for q != nil{
		head.Next = next
		next = head
		head = q
		q = head.Next
		if q == nil {
			head.Next = next
			return head
		}
	}

	return head
}
